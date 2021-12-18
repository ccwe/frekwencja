package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"os/exec"
	"runtime"

	"github.com/pinkie-py/gradr/librus"
	"golang.org/x/term"
)

type results struct {
	Attendance librus.Attendances
	GPAs       librus.GPAs
	GPA        [2]float64
}

const host = "https://frekwencja.ccwe.pl"

var Reset = "\033[0m"
var Red = "\033[31m"
var Cyan = "\033[36m"
var Gray = "\033[37m"

func init() {
	if runtime.GOOS == "windows" {
		Reset = ""
		Red = ""
		Cyan = ""
		Gray = ""
	}
}

func main() {
	fmt.Printf("Frekwencja CCWE v1.6\nCopyright (C) 2021-2022 Centrum Cyfrowego Wsparcia Edukacji\n\n")
	fmt.Printf("Wprowadź poniżej swoje dane logowania do Synergii. Gdy program pobierze dane z dziennika i obliczy wszystkie wartości, wyniki zostaną wypisane tutaj lub zostaniesz przekierowany na stronę <frekwencja.ccwe.pl>, gdzie wyświetlisz obliczone frekwencje i średnie ważone dla poszczególnych przedmiotów w tabelce (wybierzesz w trakcie).\n\n")
	fmt.Printf("Wszystkie procesy odbywają się wyłącznie na twoim urządzeniu. CCWE nie zbiera ani nie przetwarza twoich danych na żadnym etapie.\n\n")
	fmt.Printf("Uwaga! Login i hasło nie wyświetlają się podczas wpisywania.\n\n")

	running := true
	for running {
		fmt.Printf("Login: ")
		login, _ := term.ReadPassword(0)
		fmt.Printf("****\nHasło: ")
		pass, _ := term.ReadPassword(0)
		fmt.Printf("****\n\n")

		client, err := librus.New(librus.NewParams{
			Username: string(login),
			Password: string(pass),
			Debug:    true,
		})
		if err != nil {
			fmt.Printf("Coś poszło nie tak:\n  %s\n\n", err.Error())
			switch err.(type) {
			case librus.ErrInvalidGrant:
				fmt.Printf("Sprawdź swoje dane i spróbuj ponownie\n\n")
			case librus.ErrRequestInternal:
				fmt.Printf("Sprawdź czy masz połączenie z internetem i/lub ustawienia sieciowe i spróbuj ponownie\n\n")
			}
			continue
		}

		running = false

		at := client.GetAttendance()
		gr, gpa := librus.GPA(client.GetGrades())
		data := &results{at, gr, gpa}

		target := "s"
		fmt.Println("\nChcesz wyświetlić wyniki tutaj czy na stronie? Wpisz: t - tutaj, s - strona")
		fmt.Printf("(t/s): ")
		fmt.Scanf("%s", &target)
		fmt.Printf("\n")

		if target == "t" {
			var sem uint8 = 0
			fmt.Println("Który semestr? Wpisz: 1 - pierwszy, 2 - drugi, 3 - oba")
			fmt.Printf("(1/2/3): ")
			fmt.Scanf("%d", &sem)
			fmt.Printf("\n\n")
			display(data, (sem-1)%3)
			fmt.Scanf("%s", &target)
		} else {
			json_data, _ := json.Marshal(data)
			// ioutil.WriteFile("./out.ignore.json", json_data, 0666)

			link := host + "/#" + url.QueryEscape(string(compress(json_data)))
			err := openbrowser(link)
			if err != nil {
				fmt.Printf("Coś poszło nie tak:\n  %s", err.Error())
				fmt.Printf("\n%s\n\nSkopiuj powyższy link i wklej go w swojej przeglądarce, następnie naciśnij cokolwiek by zamknąć program...", link)
				fmt.Scan()
			}
		}
	}
	// json_data, _ := ioutil.ReadFile("./out.ignore.json")
	// link := host + "/#" + url.QueryEscape(string(compress(json_data)))
	// openbrowser(link)
}

func openbrowser(url string) error {
	switch runtime.GOOS {
	case "android", "ios", "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("main: failed to open the browser: unsupported platform")
	}
}

func display(r *results, s uint8) {
	fmt.Printf("Semestr: %d\n\n", s+1)
	fmt.Printf("%-35s %s  %s\n\n", "Przedmiot", "Średnia", "Frekwencja")

	for k, v := range r.Attendance {
		gpa, ok := r.GPAs[k]
		clr_at := Cyan
		if v[s].Value < 50 {
			clr_at = Red
		}
		if ok {
			clr_gpa := Cyan
			if gpa[s] < 1.75 {
				clr_gpa = Red
			}
			fmt.Printf("%-35s %s%7v %s%10v%%%s\n", k, clr_gpa, gpa[s], clr_at, v[s].Value, Reset)
		} else {
			fmt.Printf("%-35s %s%7s %s%10v%%%s\n", k, Gray, "brak", clr_at, v[s].Value, Reset)
		}
	}
	fmt.Printf("\nPrzewidywana końcowa średnia: %.2f\n", r.GPA[s])
	fmt.Println("Naciśnij Enter by wyjść...")
}

func compress(r []byte) []byte {
	data := bytes.Buffer{}
	b64 := base64.NewEncoder(base64.URLEncoding, &data)

	gz, _ := gzip.NewWriterLevel(b64, gzip.BestCompression)
	gz.Write(r)

	gz.Close()
	b64.Close()

	b64s, _ := io.ReadAll(&data)
	return b64s
}
