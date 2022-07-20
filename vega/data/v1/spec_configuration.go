package v1

func (d *DataSpecConfiguration) ToOracleSpec() *DataSpec {
	return NewOracleSpec(d.PubKeys, d.Filters)
}
