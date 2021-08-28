package atm

type Atm struct {
	dao *AtmDao
}

func newAtm(dao *AtmDao) Atm {
	return Atm{dao}
}
