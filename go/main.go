package main

import (
        "io"
        "net/http"
        "strconv"
        "strings"
)

// Fibonacci returns the first `n` Fibonacci numbers
func Fibonacci(n int) []int {
        // TODO: Fill this function out
        f := make([]int, n+1, n+2)
        if n < 2 {
            f = f[0:2]
        }
        f[0] = 0
        f[1] = 1
        for i := 2; i <= n; i++ {
            f[i] = f[i-1] + f[i-2]
        }
        return f
}

// getFibs is the response handler for the web server
func getFibs(w http.ResponseWriter, r *http.Request) {
        param, ok := r.URL.Query()["n"]
        if !ok || len(param[0]) < 1 {
                w.WriteHeader(http.StatusUnprocessableEntity)
                return
        }

        n, err := strconv.Atoi(param[0])
        if err != nil {
                w.WriteHeader(http.StatusUnprocessableEntity)
                return
        }

        nums := Fibonacci(n)
        strNums := make([]string, len(nums))

        // convert nums from int to string array
        for i, num := range nums {
                strNums[i] = strconv.Itoa(num)
        }

        finalNums := strings.Join(strNums, ", ")
        io.WriteString(w, finalNums)
        return
}

func main() {
        http.HandleFunc("/", getFibs)
        http.ListenAndServe(":80", nil)
}
