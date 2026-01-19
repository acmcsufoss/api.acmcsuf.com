package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/acmcsufoss/api.acmcsuf.com/internal/db/models"
)

type BoardService struct {
	q  *models.Queries
	db models.DBTX
}

type Position struct {
	Title string `json:"title"`
	Tier  int64  `json:"tier"`
}

type TierJSON struct {
	ID    int64 `json:"id"`
	Index int64 `json:"index"`
}

// Needed because the officers.json file stores officer and position data together
type OfficerPositions struct {
	FullName  string                `json:"fullName"`
	Picture   sql.NullString        `json:"picture"`
	Positions map[string][]Position `json:"positions"`
	Discord   sql.NullString        `json:"discord,omitempty"`
	Github    sql.NullString        `json:"github,omitempty"`
}

func main() {
	var s *BoardService
	ctx := context.Background()

	db, err := sql.Open("sqlite3", "./dev.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	/*
		var tier models.CreateTierParams
		var officer models.CreateOfficerParams
		var position models.CreatePositionParams
	*/

	// Populating tiers
	data, err := os.ReadFile("tiers.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	tiers := make(map[string]TierJSON)
	err = json.Unmarshal(data, &tiers)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	for title, t := range tiers {
		/*
			tier.Tier = t.ID
			tier.Title = title
			tier.TIndex = t.Index
			tier.Team = sql.NullString
		*/

		tier := models.CreateTierParams{
			Tier:   t.ID,
			Title:  sql.NullString{String: title, Valid: true},
			TIndex: sql.NullInt64{Int64: t.Index, Valid: true},
		}

		if _, err := s.q.CreateTier(ctx, tier); err != nil {
			log.Fatal("Error creating tier:", err)
		}
	}

	// Populating officers and positions
	data, err := os.ReadFile("officers.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var officerPositions []OfficerPositions
	err = json.Unmarshal(data, &officerPositions)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	// Splits up officerPositions into officer data and position data, then populates
	for i := range officerPositions {
		officer.Uuid = fmt.Sprintf("%04d", i+1)
		officer.FullName = officerPositions[i].FullName
		officer.Picture = officerPositions[i].Picture
		officer.Discord = officerPositions[i].Discord
		officer.Github = officerPositions[i].Github

		s.q.CreateOfficer(ctx, officer)

		position.Oid = officer.Uuid

		for semester, value := range officerPositions[i].Positions {
			position.Semester = semester

			for _, role := range value {
				position.Tier = role.Tier
				s.q.CreatePosition(ctx, position)
			}
		}
	}

	/*
		var officers []Officer
		err = json.Unmarshal(data, &officers)
		if err != nil {
			log.Fatal("Error unmarshaling JSON:", err)
		}

		data, err = os.ReadFile("officers.json")
		if err != nil {
			log.Fatal("Error reading JSON file:", err)
		}


		db, err = sql.Open("sqlite3", "./dev.db")
		if err != nil {
			log.Fatal("Error opening database:", err)
		}
		defer db.Close()


				// Create tables if they don't exist
				// DIRECT SQL CALL (will remove later)
				_, err = db.Exec(`
			        CREATE TABLE IF NOT EXISTS officer (
			            uuid CHAR(4) PRIMARY KEY,
			            full_name VARCHAR(30) NOT NULL,
			            picture VARCHAR(37),
			            github VARCHAR(64),
			            discord VARCHAR(32)
			        );

			        CREATE TABLE IF NOT EXISTS tier (
			            tier INT PRIMARY KEY,
			            title VARCHAR(40),
			            t_index INT,
			            team VARCHAR(20)
			        );

			        CREATE TABLE IF NOT EXISTS position (
			            oid CHAR(4) NOT NULL,
			            semester CHAR(3) NOT NULL,
			            tier INT NOT NULL,
			            full_name VARCHAR(30) NOT NULL,
			            title VARCHAR(40),
			            team VARCHAR(20),
			            PRIMARY KEY (oid, semester, tier),
			            CONSTRAINT fk_officers FOREIGN KEY (oid) REFERENCES officer (uuid),
			            CONSTRAINT fk_tiers FOREIGN KEY (tier) REFERENCES tier(tier)
			        );
			    `)
				if err != nil {
					log.Fatal("Error creating tables:", err)
				}

				// Prepare statements
				// DIRECT SQL CALL (will remove later)
				officerStmt, err := db.Prepare("INSERT OR IGNORE INTO officer (uuid, full_name, picture, discord) VALUES (?, ?, ?, ?)")
				if err != nil {
					log.Fatal("Error preparing officer statement:", err)
				}
				defer officerStmt.Close()

				// DIRECT SQL CALL (will remove later)
				tierStmt, err := db.Prepare("INSERT OR IGNORE INTO tier (tier, title) VALUES (?, ?)")
				if err != nil {
					log.Fatal("Error preparing tier statement:", err)
				}
				defer tierStmt.Close()

				// DIRECT SQL CALL (will remove later)
				positionStmt, err := db.Prepare("INSERT OR IGNORE INTO position (oid, semester, tier, full_name, title) VALUES (?, ?, ?, ?, ?)")
				if err != nil {
					log.Fatal("Error preparing position statement:", err)
				}
				defer positionStmt.Close()

				// Insert officers
				for i, officer := range officers {
					// Generate sequential 4-digit ID (0001, 0002, etc.)
					sequentialID := fmt.Sprintf("%04d", i+1)

					// Insert officer
					// DIRECT SQL CALL (will remove later)
					_, err = officerStmt.Exec(sequentialID, officer.FullName, officer.Picture, officer.Discord)
					if err != nil {
						log.Printf("Error inserting officer %s: %v", officer.FullName, err)
						continue
					}

					// Insert all positions the officer holds
					for semester, positions := range officer.Positions {
						for _, pos := range positions {
							// Insert tier if it doesn't exist yet
							// DIRECT SQL CALL (will remove later)
							_, err = tierStmt.Exec(pos.Tier, pos.Title)
							if err != nil {
								log.Printf("Error inserting tier %d for officer %s: %v", pos.Tier, officer.FullName, err)
								continue
							}

							// Insert position
							// DIRECT SQL CALL (will remove later)
							_, err = positionStmt.Exec(
								sequentialID,
								strings.ToUpper(semester),
								pos.Tier,
								officer.FullName,
								pos.Title,
							)
							if err != nil {
								log.Printf("Error inserting position for officer %s: %v", officer.FullName, err)
								continue
							}
						}
					}
				}
	*/

	fmt.Println("Database population completed successfully!")
	fmt.Printf("Processed %d officers\n", len(officerPositions))
}
