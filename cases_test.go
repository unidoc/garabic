package arabic

//removeHarakatTestCases contains all test cases for TestRemoveHarakat function
var removeHarakatTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "Removing Alif Khanjariyah",
		input:       "رَحْمَٰن",
		expected:    "رحمن",
	},
	{
		description: "Removing Waslah",
		input:       "ٱمْشُوا",
		expected:    "امشوا",
	},

	{
		description: "Removing all harakat 1",
		input:       "يَا أَيُّهَا الَّذِينَ آمَنُوا أَوْفُوا بِالْعُقُودِ",
		expected:    "يا أيها الذين آمنوا أوفوا بالعقود",
	},
	{
		description: "Removing all harakat 2",
		input:       "سَنواتٌ قَليلةٌ وأَدْخُلُ الجامِعَةَ، يا لها مِنْ رِحْلةٍ شاقَّةٍ طَويلَةٍ، ما أصعبَ أيامَ الدِّراسَةِ",
		expected:    "سنوات قليلة وأدخل الجامعة، يا لها من رحلة شاقة طويلة، ما أصعب أيام الدراسة",
	},
	{
		description: "Removing all harakat 3",
		input:       "إِنِّني أَشْكُرُ رَبِّي دائماً، لكنني مُنْزعجةٌ.. أَلاَ يَحِقُّ لِيَ التعبيرُ عن ضِيقِ صَدْري؟!",
		expected:    "إنني أشكر ربي دائما، لكنني منزعجة.. ألا يحق لي التعبير عن ضيق صدري؟!",
	},
}

//normalizeTestCases contains all test cases for TestNormalize function
var normalizeTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		description: "AlefHamzaAbove removal from beginning of word",
		input:       "أحمد",
		expected:    "احمد",
	},
	{
		description: "AlefHamzaAbove removal from middle of word",
		input:       "مأمون",
		expected:    "مامون",
	},
	{
		description: "AlefHamzaAbove removal from end of word",
		input:       "نمأ",
		expected:    "نما",
	},
	{
		description: "Replacing DotlessYae with Yae",
		input:       "منى",
		expected:    "مني",
	},
	{
		description: "Replacing TehMarbuta with Hae",
		input:       "مكتبة",
		expected:    "مكتبه",
	},
	{
		description: "Trimming Tatweel",
		input:       "بريـــــــــد",
		expected:    "بريد",
	},
	{
		description: "Removing Maddah",
		input:       "قُرْآن",
		expected:    "قران",
	},
	{
		description: "Normalizing Alif Waslah",
		input:       "ٱمْشُوا",
		expected:    "امشوا",
	},
	{
		description: "Removing all harakat combined with text normalization 1",
		input:       "يَا أَيُّهَا الَّذِينَ آمَنُوا أَوْفُوا بِالْعُقُودِ",
		expected:    "يا ايها الذين امنوا اوفوا بالعقود",
	},
	{
		description: "Removing all harakat combined with text normalization 2",
		input:       "سَنواتٌ قَليلةٌ وأَدْخُلُ الجامِعَةَ، يا لها مِنْ رِحْلةٍ شاقَّةٍ طَويلَةٍ، ما أصعبَ أيامَ الدِّراسَةِ",
		expected:    "سنوات قليله وادخل الجامعه، يا لها من رحله شاقه طويله، ما اصعب ايام الدراسه",
	},
	{
		description: "Removing all harakat combined with text normalization 3",
		input:       "إِنِّني أَشْكُرُ رَبِّي دائماً، لكنني مُنْزعجةٌ.. أَلاَ يَحِقُّ لِيَ التعبيرُ عن ضِيقِ صَدْري؟!",
		expected:    "انني اشكر ربي دائما، لكنني منزعجه.. الا يحق لي التعبير عن ضيق صدري؟!",
	},
}

//spellNumberTestCases contains all test cases for reading a number in arabic
var spellNumberTestCases = []struct {
	input    int
	expected string
}{
	{
		0,
		"صفر",
	},
	{
		1,
		"واحد",
	},
	{
		-1,
		"سالب واحد",
	},
	{
		2,
		"اثنان",
	},
	{
		100,
		"مئة",
	},
	{
		1000,
		"ألف",
	},
	{
		1250,
		"ألف و مئتان و خمسون",
	},
	{
		11225,
		"أحد عشر ألف و مئتان و خمسة و عشرون",
	},
	{
		100000,
		"مئة ألف",
	},
	{
		1000000,
		"مليون",
	},
	{
		-2000000,
		"سالب اثنان مليون",
	},
	{
		141592653589,
		"مئة و واحد و أربعون مليار و خمسمئة و اثنان و تسعون مليون و ستمئة و ثلاثة و خمسون ألف و خمسمئة و تسعة و ثمانون",
	},
	{
		141592653589,
		"مئة و واحد و أربعون مليار و خمسمئة و اثنان و تسعون مليون و ستمئة و ثلاثة و خمسون ألف و خمسمئة و تسعة و ثمانون",
	},
}

//spellNumberTestCases contains all test cases for reading a number in arabic
var tashkeelTestCases = []struct {
	description string
	input       string
	expected    string
}{
	{
		"Adding Kasrah after 'من'",
		"يقرأ محمد مِنَ الكتاب",
		"يقرأ محمد مِنَ الكتابِ",
	},
}
