package generator

type Generator struct {
	Length                   int    // Длина генерируемого кода
	UseLowerCaseLetters      bool   // Использовать строчные буквы
	UseUpperCaseLetters      bool   // Использовать заглавные буквы
	UseNumbers               bool   // Использовать цифры
	Symbols                  string // Символы которые будут использоваться
	UseSymbols               bool   // Использовать символы
	ExcludeSimilarCharacters bool   // Исключить похожие символы например'i', 'l', '1', 'L', 'o', 'O', '0'
	Exclude                  string // Символы, которые нужно исключить из набора символов для генерации кода
	Strict                   bool   // Должен пароль включать хотя бы по одному символу из каждой группы символов (строчные буквы, заглавные буквы, цифры, символы)}
	Beautify                 bool   // По красоте
}
