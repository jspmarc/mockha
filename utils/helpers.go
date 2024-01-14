package utils

import "database/sql"

func StrPtrToSqlNullString(str *string) sql.NullString {
	if str == nil {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}

	return sql.NullString{
		String: *str,
		Valid:  true,
	}
}

func SqlNullStringToStrPtr(nullString sql.NullString) *string {
	if nullString.Valid {
		return &nullString.String
	}

	return nil
}

func BytePtrToSqlNullByte(b *byte) sql.NullByte {
	if b == nil {
		return sql.NullByte{
			Byte:  0x0,
			Valid: false,
		}
	}

	return sql.NullByte{
		Byte:  *b,
		Valid: true,
	}
}

func SqlNullByteToBytePtr(nullString sql.NullByte) *byte {
	if nullString.Valid {
		return &nullString.Byte
	}

	return nil
}
