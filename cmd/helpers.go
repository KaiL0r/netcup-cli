package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// ======================
// Helpers
// ======================

func parseID(s string) (int, error) {
	id, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid id %q: %w", s, err)
	}
	return id, nil
}

func parseBoolArg(s string) (bool, error) {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("invalid bool %q: %w", s, err)
	}
	return b, nil
}

func printJSON(v any) error {
	out, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}
