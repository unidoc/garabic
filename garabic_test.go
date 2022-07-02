package garabic

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

const succeed = "\u2705"
const failed = "\u274C"

// TestRemoveHarakat.
func TestRemoveHarakat(t *testing.T) {

	t.Log("Given an arabic string it should be normalized")
	{
		for i, tt := range removeHarakatTestCases {
			normalized := RemoveHarakat(tt.input)
			t.Logf("\tTest: %d\t Normalizing %s", i, tt.input)
			if normalized != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be normalized to %s, got %s instead", failed, tt.description, tt.expected, normalized)
			} else {
				t.Logf("\t%s\t(%s)\tShould be normalized to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestNormalizeBigText.
func TestNormalizeBigText(t *testing.T) {
	originalArabicText, err := ioutil.ReadFile("test_data/bigText.txt")
	if err != nil {
		t.Errorf("\t%s\t Reading file failed with error:(%s)\t", failed, err)
	}
	preNormalizedArabicText, err := ioutil.ReadFile("test_data/normalizedBigText.txt")
	if err != nil {
		t.Errorf("\t%s\t Reading prenormalized file failed with error:(%s)\t", failed, err)
	}

	// Try to normalize test.
	normalized := Normalize(string(originalArabicText))
	if normalized != string(preNormalizedArabicText) {
		t.Errorf("\t%s\t Normalized text doesn't match [length of the normalized version: %d\t length of the original prenormalized version: %d\t", failed, len(normalized), len(preNormalizedArabicText))
	}
	t.Logf("\t%s\t Should normalize all text file\t", succeed)
}

// TestNormalize.
func TestNormalize(t *testing.T) {
	t.Log("Given an arabic string it should be normalized")
	{
		for i, tt := range normalizeTestCases {
			normalized := Normalize(tt.input)
			t.Logf("\tTest: %d\t Normalizing %s", i, tt.input)
			if normalized != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be normalized to %s, got %s instead", failed, tt.description, tt.expected, normalized)
			} else {
				t.Logf("\t%s\t(%s)\tShould be normalized to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestDeleteRune.
func TestDeleteRune(t *testing.T) {
	testCases := []struct {
		description string
		input       []rune
		index       int
		expected    []rune
	}{
		{
			description: "Deleting rune with index 0 that exists in the array",
			input:       []rune{'م', 'ح'},
			index:       0,
			expected:    []rune{'ح'},
		},
		{
			description: "Deleting rune with index 3 that exists in the array",
			input:       []rune{'م', 'ا', 'د', 'ة', 'ن'},
			index:       3,
			expected:    []rune{'م', 'ا', 'د', 'ن'},
		},
		{
			description: "Deleting rune with index 0 that doesn't exist in the array",
			input:       []rune{},
			index:       0,
			expected:    []rune{},
		},
		{
			description: "Deleting rune with index 0 that's more than array length",
			input:       []rune{'أ', 'ب'},
			index:       10,
			expected:    []rune{'أ', 'ب'},
		},
	}

	t.Log("Given a slice of runes and a position of a rune, it should be deleted from the slice while keeping order")
	{
		for i, tt := range testCases {
			input := tt.input
			input = deleteRune(input, tt.index)
			t.Logf("\tTest: %d\t Deleting rune %d from %v", i, tt.index, tt.input)
			if !reflect.DeepEqual(input, tt.expected) {
				t.Errorf(
					"\t%s\t(%s)\tShould be %v (len %d, cap %d), got %v (len %d, cap %d) instead",
					failed, tt.description, tt.expected, len(tt.expected), cap(tt.expected), input, len(input), cap(input),
				)
			} else {
				t.Logf("\t%s\t(%s)\tShould be %v", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestSpellNumber.
func TestSpellNumber(t *testing.T) {
	t.Log("Given a number it should be return readable string of it in arabic")
	{
		for i, tt := range spellNumberTestCases {
			textOfNum := SpellNumber(tt.input)
			t.Logf("\tTest: %d\t Spelling Number %d", i, tt.input)
			if textOfNum != tt.expected {
				t.Errorf("\t%s\t\tShould be converted to %s, got %s instead", failed, tt.expected, textOfNum)
			} else {
				t.Logf("\t%s\t\tShould be converted to %s", succeed, tt.expected)
			}
		}
	}
}

// TestTashkeel.
func TestTashkeel(t *testing.T) {
	t.Log("Given an arabic string, diacritics should be added correctly")
	{
		for i, tt := range tashkeelTestCases {
			withDiacritics := Tashkeel(tt.input)
			t.Logf("\tTest: %d\t Adding diacritics to %s", i, tt.input)
			if withDiacritics != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be updated to %s, got %s instead", failed, tt.description, tt.expected, withDiacritics)
			} else {
				t.Logf("\t%s\t(%s)\tShould be updated to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

// TestShape.
func TestShape(t *testing.T) {
	t.Log("Given an arabic string, shaping will be fixed for rendering")
	{
		for i, tt := range shapingTestCases {
			shapedArabicText := Shape(tt.input)
			t.Logf("\tTest: %d\t Shaping: %s", i, tt.input)
			if shapedArabicText != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould be updated to \n'%s'\n, got \n\"%s\"\n instead", failed, tt.description, tt.expected, shapedArabicText)
			} else {
				t.Logf("\t%s\t(%s)\tShould be updated to %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

func TestIsArabicLetter(t *testing.T) {
	t.Log("Given a letter, check if it's an arabic letter")
	{
		for i, tt := range arabicLetterTestCases {
			t.Logf("\tTest: %d\t Checking if letter %s is arabic", i, string(tt.input))
			if IsArabicLetter(tt.input) != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould return %t, got %t instead", failed, tt.description, tt.expected, IsArabicLetter(tt.input))
			} else {
				t.Logf("\t%s\t(%s)\tShould be %t", succeed, tt.description, tt.expected)
			}
		}
	}
}

func TestIsArabic(t *testing.T) {
	t.Log("Given a text, check if it's arabic")
	{
		for i, tt := range arabicTextTestCases {
			t.Logf("\tTest: %d\t Checking if letter %s is arabic", i, string(tt.input))
			if IsArabic(tt.input) != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould return %t, got %t instead", failed, tt.description, tt.expected, IsArabic(tt.input))
			} else {
				t.Logf("\t%s\t(%s)\tShould be %t", succeed, tt.description, tt.expected)
			}
		}
	}
}

func TestToArabicDigits(t *testing.T) {
	t.Log("Given a text, convert all english digits in it into arabic digits")
	{
		for i, tt := range arabicNumbersTestCases {
			t.Logf("\tTest: %d\t Checking if letter %s is arabic", i, string(tt.input))
			if ToArabicDigits(tt.input) != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould return %s, got %s instead", failed, tt.description, tt.expected, ToArabicDigits(tt.input))
			} else {
				t.Logf("\t%s\t(%s)\tShould be %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

func TestToEnglishDigits(t *testing.T) {
	t.Log("Given a text, convert all arabic digits in it into english digits")
	{
		for i, tt := range englishNumbersTestCases {
			t.Logf("\tTest: %d\t Checking if letter %s is arabic", i, string(tt.input))
			if ToEnglishDigits(tt.input) != tt.expected {
				t.Errorf("\t%s\t(%s)\tShould return %s, got %s instead", failed, tt.description, tt.expected, ToEnglishDigits(tt.input))
			} else {
				t.Logf("\t%s\t(%s)\tShould be %s", succeed, tt.description, tt.expected)
			}
		}
	}
}

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range normalizeTestCases {
			Normalize(c.input)
		}
	}
}

func BenchmarkRemoveHarakat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range removeHarakatTestCases {
			RemoveHarakat(c.input)
		}
	}
}

func BenchmarkNormalizeBigText(b *testing.B) {
	originalArabicText, err := ioutil.ReadFile("test_data/bigText.txt")
	if err != nil {
		b.Errorf("\t%s\t Reading file failed with error:(%s)\t", failed, err)
	}
	for i := 0; i < b.N; i++ {
		Normalize(string(originalArabicText))
	}
}

func ExampleNormalize() {
	normalized := Normalize("أحمد")
	fmt.Println(normalized)
	// Output:
	// احمد
}

func ExampleRemoveHarakat() {
	normalized := RemoveHarakat("سَنواتٌ")
	fmt.Println(normalized)
	// Output:
	// سنوات
}

func ExampleSpellNumber() {
	numberInWords := SpellNumber(100)
	fmt.Println(numberInWords)
	// Output:
	// مئة
}

func ExampleIsArabicLetter() {
	fmt.Println(IsArabicLetter('ص'))
	// Output:
	// true
}

func ExampleToArabicDigits() {
	fmt.Println(ToArabicDigits("عام 2021"))
	// Output:
	// عام ٢٠٢١
}

func ExampleToEnglishDigits() {
	fmt.Println(ToEnglishDigits("عام ٢٠٢١"))
	// Output:
	// عام 2021
}
