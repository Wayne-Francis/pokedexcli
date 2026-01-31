package main
import ("strings")

func cleanInput(text string) []string {
	lower_text := strings.ToLower(text)
	clean_output := strings.Fields(lower_text)
return clean_output
}