// go build ex01.go
// powstaje plik wykonywalny w tym samym folderze, co źródło

// go get - możliwość ściągania kody z repozytorium

// pakiety ~ biblioteki, moduły
// pakiet main definiuje samodzielny wykonywalny program

// zmiana nazwy pakietu - nie mamy wtedy kilku źródeł deklarujących funkcję main
// ciekawe - wystarczy zmienić nazwę pakietu z main tylko w pierwszym pliku i go-plus nie narzeka

package ex01

// Go nie pozwala na zbędne importy

import "fmt"

// można użyć średnika dla oddzielenia kilku wyrażeń w jednej linii
// po niektórych znakach nowe linie są konwertowane na średniki; klamra otwierająca musi być w jednej linii z nagłówkiem funkcji; w wyrażeniu x + y nowa linia może być po +, ale nie przed +

// gofmt - narzędzie do formatowania kodu w standardowym formacie
// gofmt fmt - formatowanie w pakiecie lub w tym samym folderze (domyślnie)

// goimports - zarządzanie deklaracjami import, nie jest to część Go, można ściągnąć: go get golang.org/x/tools/cmd/goimports

func main() {
	fmt.Println("Hello, Go!")
}
