package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

import "github.com/acmcsufoss/api.acmcsuf.com/internal/api/store/dbmodels"

type Role struct {
	Title string `json:"title"`
	Tier  int64  `json:"tier"`
}

type Tier struct {
	ID    int `json:"id"`
	Index int `json:"index"`
}

type Member struct {
	FullName  string            `json:"fullName"`
	Picture   string            `json:"picture"`
	Positions map[string][]Role `json:"positions"`
	Discord   string            `json:"discord"`
	Github    string            `json:"github"`
}

func toNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

// Helper to convert a numeric int into a nullable database NullString parameter
func intToNullString(val int) sql.NullString {
	return sql.NullString{
		String: fmt.Sprintf("%d", val),
		Valid:  true,
	}
}

// Helper to convert an int variable into a database safe sql.NullInt64 parameter
func toNullInt64(val int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(val),
		Valid: true,
	}
}

func generateUUID() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Set version to 4 (pseudo-random)
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	// Set variant to RFC 4122
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	return fmt.Sprintf("%x-%x-%x-%x-%x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16]), nil
}

// Helper function to fetch remote JSON files over HTTP
func fetchJSON(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func main() {
	ctx := context.Background()

	officersURL := "https://raw.githubusercontent.com/acmcsufoss/acmcsuf.com/main/src/lib/public/board/data/officers.json"
	tiersURL := "https://raw.githubusercontent.com/acmcsufoss/acmcsuf.com/main/src/lib/public/board/data/tiers.json"

	fileBytes, err := fetchJSON(officersURL)
	if err != nil {
		log.Fatalf("Error reading officers JSON from URL: %v", err)
	}

	fBytes, err := fetchJSON(tiersURL)
	if err != nil {
		log.Fatalf("Error reading tiers JSON from URL: %v", err)
	}

	var tiers map[string]Tier
	if err := json.Unmarshal(fBytes, &tiers); err != nil {
		log.Fatalf("Failed to decode JSON contents: %v", err)
	}

	var members []Member
	if err := json.Unmarshal(fileBytes, &members); err != nil {
		log.Fatalf("Failed to decode JSON contents: %v", err)
	}

	db, err := sql.Open("sqlite3", "local.db")
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer db.Close()

	queries := dbmodels.New(db)

	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to begin db transaction: %v", err)
	}
	defer tx.Rollback()

	qtx := queries.WithTx(tx)

	for _, m := range members {
		officerUUID, err := generateUUID()
		if err != nil {
			log.Fatalf("Failed to generate native cryptographic UUID: %v", err)
		}

		officerParams := dbmodels.CreateOfficerParams{
			Uuid:     officerUUID,
			FullName: m.FullName,
			Picture:  toNullString(m.Picture),
			Github:   toNullString(m.Github),
			Discord:  toNullString(m.Discord),
		}

		_, err = qtx.CreateOfficer(ctx, officerParams)
		if err != nil {
			log.Fatalf("Failed creating officer %s: %v", m.FullName, err)
		}

		for semester, roles := range m.Positions {
			for _, r := range roles {
				positionParams := dbmodels.CreatePositionParams{
					OfficerID: officerUUID,
					Semester:  semester,
					Tier:      r.Tier,
					FullName:  m.FullName,
					Title:     toNullString(r.Title),
					Team:      sql.NullString{Valid: false},
				}

				_, err := qtx.CreatePosition(ctx, positionParams)
				if err != nil {
					log.Fatalf("Failed creating position for officer %s: %v", m.FullName, err)
				}
			}
		}
	}

	for tierName, t := range tiers {
		tierParams := dbmodels.CreateTierParams{
			Tier:   int64(t.ID),
			Title:  toNullString(tierName),
			TIndex: toNullInt64(t.Index),
			Team:   sql.NullString{Valid: false},
		}

		_, err := qtx.CreateTier(ctx, tierParams)
		if err != nil {
			log.Fatalf("Failed creating tier %s: %v", tierName, err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatalf("Transaction commit failed: %v", err)
	}

	log.Println("Successfully integrated JSON into SQLite via sqlc from remote URLs!")
}