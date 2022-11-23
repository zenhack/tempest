package database

func (tx Tx) AddPackage(pkgId string) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO packages(id) VALUES (?)`,
		pkgId,
	)
	return err
}

type NewGrain struct {
	GrainId string
	PkgId   string
	OwnerId string
	Title   string
}

type NewAccount struct {
	Id      string
	IsAdmin bool
	Profile struct {
		DisplayName     string
		PreferredHandle string
	}
}

type NewCredential struct {
	AccountId string
	Login     bool
	Type      string
	ScopedId  string
}

func (tx Tx) AddAccount(a NewAccount) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO accounts
			( id
			, isAdmin
			, profile_display_name
			, profile_preferred_handle
			)
			VALUES (?, ?, ?, ?)`,
		a.Id,
		a.IsAdmin,
		a.Profile.DisplayName,
		a.Profile.PreferredHandle,
	)
	return err
}

func (tx Tx) AddCredential(c NewCredential) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO credentials
			( account_id
			, login
			, type
			, scoped_id
			)
			VALUES (?, ?, ?, ?)`,
		c.AccountId,
		c.Login,
		c.Type,
		c.ScopedId,
	)
	return err
}

func (tx Tx) AddGrain(g NewGrain) error {
	_, err := tx.sqlTx.Exec(
		`INSERT INTO grains(id, packageId, title, ownerId) VALUES (?, ?, ?, ?)`,
		g.GrainId, g.PkgId, g.Title, g.OwnerId,
	)
	return err
}

// Returns the package id for the specified grain
func (tx Tx) GetGrainPackageId(grainId string) (string, error) {
	row := tx.sqlTx.QueryRow("SELECT packageId FROM grains WHERE id = ?", grainId)
	var result string
	err := row.Scan(&result)
	return result, err
}
