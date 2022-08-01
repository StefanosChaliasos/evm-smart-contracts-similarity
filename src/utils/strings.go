package utils

func FindIdentical(mapOfStrings map[string]string) map[string][]string {
    res := make(map[string][]string)
    for k, v := range mapOfStrings {
        res[v] = append(res[v], k)
    }
    return res
}
