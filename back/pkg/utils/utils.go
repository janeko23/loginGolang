package utils

//NumIncluded devuelve true si el elemento e existe en el slice s
func NumIncluded (elem int, s []int) bool {

	if len(s) == 0 {
		return false
	}
	for _, e := range s {
		if e == elem {
			return true;
		}
	}
	return false
}

const FromImport = "import"
const FromFormInterno = "formInterno"
const FromFormExterno = "formExterno"
const FromRequest = "request"
