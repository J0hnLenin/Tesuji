package pgstorage

func (s *PGstorage) bucketByGameID(id int64) bucketNum {
	return bucketNum(id % int64(s.bucketQuantity))
}