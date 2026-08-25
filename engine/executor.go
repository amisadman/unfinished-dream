package engine

func Execute(stmt *Statement) string {
	switch stmt.Type {
	case "SELECT":
		return "(0 rows returned) - The void is empty."
	case "INSERT":
		return "1 row inserted (into the void)."
	case "CREATE":
		return "Table created (in my imagination)."
	default:
		return "Nothing happened."
	}
}
