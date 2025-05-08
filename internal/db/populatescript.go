package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Officer struct {
	FullName  string            `json:"fullName"`
	Picture   string            `json:"picture"`
	Positions map[string][]struct {
		Title string `json:"title"`
		Tier  int    `json:"tier"`
	} `json:"positions"`
	Discord string `json:"discord,omitempty"`
}

func main() {
	//Read JSON
	data, err := os.ReadFile("officers.json")
	if err != nil {
		log.Fatal("Error reading JSON file:", err)
	}

	var officers []Officer
	err = json.Unmarshal(data, &officers)
	if err != nil {
		log.Fatal("Error unmarshaling JSON:", err)
	}

	// Open SQLite database
	db, err := sql.Open("sqlite3", "./officers.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create tables if they don't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS officers (
			uuid CHAR(4) PRIMARY KEY,
			full_name VARCHAR(30) NOT NULL,
			picture VARCHAR(37),
			github VARCHAR(64),
			discord VARCHAR(32)
		);
		
		CREATE TABLE IF NOT EXISTS tiers (
			tier INT PRIMARY KEY,
			title VARCHAR(40),
			t_index INT,
			team VARCHAR(20)
		);
		
		CREATE TABLE IF NOT EXISTS positions (
			oid CHAR(4) NOT NULL,
			semester CHAR(3) NOT NULL,
			tier INT NOT NULL,
			PRIMARY KEY (oid, semester, tier),
			CONSTRAINT fk_officers FOREIGN KEY (oid) REFERENCES officers (uuid),
			CONSTRAINT fk_tiers FOREIGN KEY (tier) REFERENCES tiers(tier)
		);
	`)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}

	// Prepare statements for inserting data
	officerStmt, err := db.Prepare("INSERT OR IGNORE INTO officers (uuid, full_name, picture, discord) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing officer statement:", err)
	}
	defer officerStmt.Close()

	tierStmt, err := db.Prepare("INSERT OR IGNORE INTO tiers (tier, title) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Error preparing tier statement:", err)
	}
	defer tierStmt.Close()

	positionStmt, err := db.Prepare("INSERT OR IGNORE INTO positions (oid, semester, tier) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing position statement:", err)
	}
	defer positionStmt.Close()

	// Process each officer with sequential ID
	for i, officer := range officers {
		// Generate sequential 4-digit ID (0001, 0002, etc.)
		sequentialID := fmt.Sprintf("%04d", i+1)

		// Insert officer
		_, err = officerStmt.Exec(sequentialID, officer.FullName, officer.Picture, officer.Discord)
		if err != nil {
			log.Printf("Error inserting officer %s: %v", officer.FullName, err)
			continue
		}

		// Process positions
		for semester, positions := range officer.Positions {
			for _, pos := range positions {
				// Insert tier if it doesn't exist
				_, err = tierStmt.Exec(pos.Tier, pos.Title)
				if err != nil {
					log.Printf("Error inserting tier %d for officer %s: %v", pos.Tier, officer.FullName, err)
					continue
				}

				// Insert position
				_, err = positionStmt.Exec(sequentialID, strings.ToUpper(semester), pos.Tier)
				if err != nil {
					log.Printf("Error inserting position for officer %s: %v", officer.FullName, err)
					continue
				}
			}
		}
	}

	fmt.Println("Database population completed successfully!")
	fmt.Printf("Processed %d officers\n", len(officers))
	defer db.Close()  // ← This should appear right after sql.Open()
}