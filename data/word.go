package data

import (
	"sort"
)

var WordKeys []string

func init() {
	// Loop through Word and put togther a list of keys
	for key := range Word {
		WordKeys = append(WordKeys, key)
	}

	// Sort the keys
	sort.Strings(WordKeys)
}

// Word consists of common english words
var Word = map[string][]string{
	// Nouns
	"noun_common":            {"time", "person", "year", "way", "day", "thing", "man", "world", "life", "hand", "part", "child", "eye", "woman", "place", "work", "week", "case", "point", "government", "company", "number", "group", "problem", "fact"},
	"noun_concrete":          {"apple", "air", "conditioner", "airport", "ambulance", "aircraft", "apartment", "arrow", "antlers", "apro", "alligator", "architect", "ankle", "armchair", "aunt", "ball", "bermudas", "beans", "balloon", "bear", "blouse", "bed", "bow", "bread", "black", "board", "bones", "bill", "bitterness", "boxers", "belt", "brain", "buffalo", "bird", "baby", "book", "back", "butter", "bulb", "buckles", "bat", "bank", "bag", "bra", "boots", "blazer", "bikini", "bookcase", "bookstore", "bus stop", "brass", "brother", "boy", "blender", "bucket", "bakery", "bow", "bridge", "boat", "car", "cow", "cap", "cooker", "cheeks", "cheese", "credenza", "carpet", "crow", "crest", "chest", "chair", "candy", "cabinet", "cat", "coffee", "children", "cookware", "chaise longue", "chicken", "casino", "cabin", "castle", "church", "cafe", "cinema", "choker", "cravat", "cane", "costume", "cardigan", "chocolate", "crib", "couch", "cello", "cashier", "composer", "cave", "country", "computer", "canoe", "clock", "dog", "deer", "donkey", "desk", "desktop", "dress", "dolphin", "doctor", "dentist", "drum", "dresser", "designer", "detective", "daughter", "egg", "elephant", "earrings", "ears", "eyes", "estate", "finger", "fox", "frock", "frog", "fan", "freezer", "fish", "film", "foot", "flag", "factory", "father", "farm", "forest", "flower", "fruit", "fork", "grapes", "goat", "gown", "garlic", "ginger", "giraffe", "gauva", "grains", "gas station", "garage", "gloves", "glasses", "gift", "galaxy", "guitar", "grandmother", "grandfather", "governor", "girl", "guest", "hamburger", "hand", "head", "hair", "heart", "house", "horse", "hen", "horn", "hat", "hammer", "hostel", "hospital", "hotel", "heels", "herbs", "host", "jacket", "jersey", "jewelry", "jaw", "jumper", "judge", "juicer", "keyboard", "kid", "kangaroo", "koala", "knife", "lemon", "lion", "leggings", "leg", "laptop", "library", "lamb", "london", "lips", "lung", "lighter", "luggage", "lamp", "lawyer", "mouse", "monkey", "mouth", "mango", "mobile", "milk", "music", "mirror", "musician", "mother", "man", "model", "mall", "museum", "market", "moonlight", "medicine", "microscope", "newspaper", "nose", "notebook", "neck", "noodles", "nurse", "necklace", "noise", "ocean", "ostrich", "oil", "orange", "onion", "oven", "owl", "paper", "panda", "pants", "palm", "pasta", "pumpkin", "pharmacist", "potato", "parfume", "panther", "pad", "pencil", "pipe", "police", "pen", "pharmacy", "police station", "parrot", "plane", "pigeon", "phone", "peacock", "pencil", "pig", "pouch", "pagoda", "pyramid", "purse", "pancake", "popcorn", "piano", "physician", "photographer", "professor", "painter", "park", "plant", "parfume", "radio", "razor", "ribs", "rainbow", "ring", "rabbit", "rice", "refrigerator", "remote", "restaurant", "road", "surgeon", "scale", "shampoo", "sink", "salt", "shark", "sandals", "shoulder", "spoon", "soap", "sand", "sheep", "sari", "stomach", "stairs", "soup", "shoes", "scissors", "sparrow", "shirt", "suitcase", "stove", "stairs", "snowman", "shower", "swan", "suit", "sweater", "smoke", "skirt", "sofa", "socks", "stadium", "skyscraper", "school", "sunglasses", "sandals", "slippers", "shorts", "sandwich", "strawberry", "spaghetti", "shrimp", "saxophone", "sister", "son", "singer", "senator", "street", "supermarket", "swimming pool", "star", "sky", "sun", "spoon", "ship", "smile", "table", "turkey", "tie", "toes", "truck", "train", "taxi", "tiger", "trousers", "tongue", "television", "teacher", "turtle", "tablet", "train station", "toothpaste", "tail", "theater", "trench coat", "tea", "tomato", "teen", "tunnel", "temple", "town", "toothbrush", "tree", "toy", "tissue", "telephone", "underwear", "uncle", "umbrella", "vest", "voice", "veterinarian", "villa", "violin", "village", "vehicle", "vase", "wallet", "wolf", "waist", "wrist", "water melon", "whale", "water", "wings", "whisker", "watch", "woman", "washing machine", "wheelchair", "waiter", "wound", "xylophone", "zebra", "zoo"},
	"noun_abstract":          {"fiction", "horror", "dream", "luck", "movement", "right", "clarity", "joy", "care", "trend", "belief", "sorrow", "joy", "failure", "slavery", "riches", "fashion", "envy", "success", "fear", "union", "luxury", "freedom", "generosity", "wit", "peace", "hatred", "thrill", "brilliance", "care", "wealth", "religion", "divorce", "goal", "stupidity", "friendship", "goodness", "rhythm", "timing", "infancy", "disregard", "riches", "appetite", "loneliness", "pleasure", "love", "beauty", "annoyance", "kindness", "nap", "gain", "talent", "religion", "lie", "truth", "solitude", "justice", "bravery", "calm", "childhood", "confusion", "ability", "loss", "thought", "growth", "cleverness", "anger", "horror", "marriage", "delay", "philosophy", "generation", "wisdom", "dishonesty", "happiness", "coldness", "poverty", "brilliance", "luxuty", "sleep", "awareness", "idea", "disregard", "slavery", "growth", "company", "irritation", "advantage", "mercy", "speed", "pain", "gossip", "crime", "comfort", "frailty", "life", "patience", "omen", "deceit", "elegance"},
	"noun_collective_people": {"band", "troupe", "dynasty", "group", "bevy", "staff", "crowd", "party", "board", "regiment", "crew", "tribe", "body", "patrol", "congregation", "pack", "bunch", "company", "team", "mob", "caravan", "line", "troop", "choir", "host", "posse", "class", "gang", "horde"},
	"noun_collective_animal": {"cackle", "mustering", "mob", "wisp", "pod", "bale", "murder", "muster", "brace", "exaltation", "party", "flock", "cast", "sedge", "stand", "scold", "team", "covey", "trip", "army", "school", "nest", "leap", "host", "troop"},
	"noun_collective_thing":  {"wad", "pair", "album", "string", "anthology", "reel", "outfit", "fleet", "comb", "archipelago", "quiver", "bale", "packet", "hedge", "basket", "orchard", "batch", "library", "battery", "set", "harvest", "block", "forest", "book", "group", "bouquet", "collection", "bowl", "stack", "bunch", "hand", "bundle", "catalog", "shower", "ream", "chest", "heap", "range", "cluster", "pack", "hail", "cloud", "galaxy", "sheaf", "clump"},
	"noun_countable":         {"camp", "hospital", "shirt", "sock", "plant", "cup", "fork", "spoon", "plate", "straw", "town", "box", "bird", "father", "answer", "egg", "purse", "mirror", "mistake", "toilet", "toothbrush", "shower", "towel", "pool", "corner", "card", "lawn", "city", "egg", "yard", "burger", "kilometer", "mile", "father", "film", "actor", "issue", "machine", "liter", "room", "station", "journey", "castle", "hour", "finger", "boy", "book", "year", "second", "son", "month", "group", "hall", "cat", "week", "picture", "day", "village", "effect", "baby", "weekend", "class", "meal", "river", "grade", "bush", "desk", "stream", "method", "brother", "sister", "factory", "aunt", "bush", "program", "uncle", "ball", "cousin", "wall", "grandmother", "cup", "grandfather", "week", "school", "shirt", "child", "king", "road", "judge", "bridge", "car", "line", "book", "eye", "teacher", "foot", "party", "face", "day", "chest", "handle", "week", "hotel", "eye", "animal", "doctor", "adult", "village", "key", "bird", "bank", "program", "idea", "gun", "card", "brother", "dress", "room", "door", "mouth", "club", "game", "ring", "project", "sister", "road", "coat", "account", "group", "cigarette", "farm", "river", "college", "computer", "walk", "corner", "cat", "head", "street", "election", "country", "chair", "crowd", "cup", "plant", "farm", "handle", "model", "book", "message", "battle", "pen", "pencil", "elephant", "carrot", "onion", "garden", "country", "engine", "bill", "apple", "noun", "club", "crowd", "window", "field", "friend", "verb", "class", "flower", "seed", "lake", "plant", "animal", "ocean", "whale", "fish", "stream", "cloud", "couch", "steak", "problem", "light", "door", "room", "painting", "shop", "apartment", "candle", "adult", "building", "plan", "page", "ball", "game", "animal", "apartment", "box", "thought", "walk", "lady", "bottle", "article", "game", "kettle", "car", "house", "hoses", "orange", "phone", "app", "window", "door", "dollar", "foot", "cent", "library", "cat", "bed", "pound", "gate", "tomatoes", "gun", "holiday", "woman", "job", "shock", "salary", "tax", "coat", "scooter", "dog", "problem", "field", "answer", "ear", "camp", "case", "road", "woman", "product", "bridge", "man", "dream", "idea", "scheme", "invention", "cigarette", "mother", "friend", "chapter", "computer", "dream", "father", "child", "motor", "deskpath", "factory", "park", "newspaper", "hat", "dream", "table", "kitchen", "student", "captain", "doctor", "bus", "neck", "class", "list", "member", "chest", "valley", "product", "horse", "captain", "star", "hour", "page", "bus", "girl", "month", "child", "house", "boy", "bill", "kitchen", "chapter", "boat", "hand", "dress", "table", "wall", "chair", "train", "minute", "magazine", "bus", "party", "bird", "lake", "job", "nation", "bike", "election", "hand", "box", "beach", "address", "project", "task", "park", "face", "college", "bell", "plane", "store", "hall", "accident", "daughter", "ship", "candy", "smile", "city", "island", "case", "spot", "film", "husband", "artist", "tour", "bag", "boat", "driver", "office", "chair", "path", "dog", "bag", "finger", "apartment", "garden", "heart", "year", "engine", "girl", "day", "castle", "plane", "ring", "brother", "edge", "picture", "meeting", "tent", "dog", "hat", "head", "bottle", "hill"},
	"noun_uncountable":       {"accommodation", "advertising", "air", "aid", "advice", "anger", "art", "assistance", "bread", "business", "butter", "calm", "cash", "chaos", "cheese", "childhood", "clothing", "coffee", "content", "corruption", "courage", "currency", "damage", "danger", "darkness", "data", "determination", "economics", "education", "electricity", "employment", "energy", "entertainment", "enthusiasm", "equipment", "evidence", "failure", "fame", "fire", "flour", "food", "freedom", "friendship", "fuel", "furniture", "fun", "genetics", "gold", "grammar", "guilt", "hair", "happiness", "harm", "health", "heat", "help", "homework", "honesty", "hospitality", "housework", "humour", "imagination", "importance", "information", "innocence", "intelligence", "jealousy", "juice", "justice", "kindness", "knowledge", "labour", "lack", "laughter", "leisure", "literature", "litter", "logic", "love", "luck", "magic", "management", "metal", "milk", "money", "motherhood", "motivation", "music", "nature", "news", "nutrition", "obesity", "oil", "old age", "oxygen", "paper", "patience", "permission", "pollution", "poverty", "power", "pride", "production", "progress", "pronunciation", "publicity", "punctuation", "quality", "quantity", "racism", "rain", "relaxation", "research", "respect", "rice", "room (space)", "rubbish", "safety", "salt", "sand", "seafood", "shopping", "silence", "smoke", "snow", "software", "soup", "speed", "spelling", "stress", "sugar", "sunshine", "tea", "tennis", "time", "tolerance", "trade", "traffic", "transportation", "travel", "trust", "understanding", "unemployment", "usage", "violence", "vision", "warmth", "water", "wealth", "weather", "weight", "welfare", "wheat", "width", "wildlife", "wisdom", "wood", "work", "yoga", "youth"},
	//"noun_proper":            {}, // This refers to an actual person(John Doe), place(Chipotle, Tennessee)

	// Verbs
	"verb_action":  {"ride", "sit", "stand", "fight", "laugh", "read", "play", "listen", "cry", "think", "sing", "watch", "dance", "turn", "win", "fly", "cut", "throw", "sleep", "close", "open", "write", "give", "jump", "eat", "drink", "cook", "wash", "wait", "climb", "talk", "crawl", "dream", "dig", "clap", "knit", "sew", "smell", "kiss", "hug", "snore", "bathe", "bow", "paint", "dive", "ski", "stack", "buy", "shake"},
	"verb_linking": {"am", "is", "was", "are", "were", "being", "been", "be", "have", "has", "had", "do", "does", "did", "shall", "will", "should", "would", "may", "might", "must", "can", "could"},
	"verb_helping": {"is", "can", "be", "do", "may", "had", "should", "was", "has", "could", "are", "will", "been", "did", "might", "were", "does", "must", "have", "would", "am", "shall", "being"},

	// Adverbs
	"adverb_manner":               {"accidentally", "angrily", "anxiously", "awkwardly", "badly", "beautifully", "blindly", "boldly", "bravely", "brightly", "busily", "calmly", "carefully", "carelessly", "cautiously", "cheerfully", "clearly", "closely", "correctly", "courageously", "cruelly", "daringly", "deliberately", "doubtfully", "eagerly", "easily", "elegantly", "enormously", "enthusiastically", "equally", "eventually", "exactly", "faithfully", "fast", "fatally", "fiercely", "fondly", "foolishly", "fortunately", "frankly", "frantically", "generously", "gently", "gladly", "gracefully", "greedily", "happily", "hard", "hastily", "healthily", "honestly", "hungrily", "hurriedly", "inadequately", "ingeniously", "innocently", "inquisitively", "irritably", "joyously", "justly", "kindly", "lazily", "loosely", "loudly", "madly", "mortally", "mysteriously", "neatly", "nervously", "noisily", "obediently", "openly", "painfully", "patiently", "perfectly", "politely", "poorly", "powerfully", "promptly", "punctually", "quickly", "quietly", "rapidly", "rarely", "really", "recklessly", "regularly", "reluctantly", "repeatedly", "rightfully", "roughly", "rudely", "sadly", "safely", "selfishly", "sensibly", "seriously", "sharply", "shyly", "silently", "sleepily", "slowly", "smoothly", "so", "softly", "solemnly", "speedily", "stealthily", "sternly", "straight", "stupidly", "successfully", "suddenly", "suspiciously", "swiftly", "tenderly", "tensely", "thoughtfully", "tightly", "truthfully", "unexpectedly", "victoriously", "violently", "vivaciously", "warmly", "weakly", "wearily", "well", "wildly", "wisely"},
	"adverb_degree":               {"almost", "absolutely", "awfully", "badly", "barely", "completely", "decidedly", "deeply", "enough", "enormously", "entirely", "extremely", "fairly", "far", "fully", "greatly", "hardly", "highly", "how", "incredibly", "indeed", "intensely", "just", "least", "less", "little", "lots", "most", "much", "nearly", "perfectly", "positively", "practically", "pretty", "purely", "quite", "rather", "really", "scarcely", "simply", "so", "somewhat", "strongly", "terribly", "thoroughly", "too", "totally", "utterly", "very", "virtually", "well"},
	"adverb_place":                {"about", "above", "abroad", "anywhere", "away", "back", "backwards", "behind", "below", "down", "downstairs", "east", "elsewhere", "far", "here", "in", "indoors", "inside", "near", "nearby", "off", "on", "out", "outside", "over", "there", "towards", "under", "up", "upstairs", "where"},
	"adverb_time_definite":        {"now", "then", "today", "tomorrow", "tonight", "yesterday"},
	"adverb_time_indefinite":      {"already", "before", "early", "earlier", "eventually", "finally", "first", "formerly", "just", "last", "late", "later", "lately", "next", "previously", "recently", "since", "soon", "still", "yet"},
	"adverb_frequency_definite":   {"annually", "daily", "fortnightly", "hourly", "monthly", "nightly", "quarterly", "weekly", "yearly"},
	"adverb_frequency_indefinite": {"always", "constantly", "ever", "frequently", "generally", "infrequently", "never", "normally", "occasionally", "often", "rarely", "regularly", "seldom", "sometimes", "regularly", "usually"},

	// Prepositions
	"preposition_simple":   {"at", "by", "as", "but", "from", "for", "into", "in", "than", "of", "off", "on", "out", "over", "till", "to", "up", "upon", "with", "under", "down"},
	"preposition_double":   {"outside of", "out of", "upon", "within", "inside", "without", "onto", "from behind", "because of", "out of", "throughout", "up to", "before", "due to", "according to", "from beneath", "next to", "from above"},
	"preposition_compound": {"according to", "as to", "onto", "across", "after", "beyond", "without", "opposite to", "away from", "aside from", "in favor of", "in front of", "because of", "as for", "near to", "behind", "along", "outside", "on account of", "on behalf of", "but for", "ahead of", "close to", "despite", "depending on", "due to", "in addition to", "next to", "in between", "in case of", "owing to", "along with", "around", "between", "apart from", "in return for", "out of", "instead of", "outside of", "other than", "together with", "up to", "above", "about"},

	// Adjectives
	"adjective_descriptive":   {"adorable", "adventurous", "agreeable", "alive", "aloof", "amused", "angry", "annoying", "anxious", "arrogant", "ashamed", "attractive", "auspicious", "awful", "bad", "beautiful", "black", "blue", "blushing", "bored", "brave", "bright", "brown", "busy", "calm", "careful", "cautious", "charming", "cheerful", "clean", "clear", "clever", "clumsy", "colorful", "comfortable", "concerning", "condemned", "confusing", "cooperative", "courageous", "creepy", "crowded", "cruel", "curios", "cute", "dangerous", "dark", "defiant", "delightful", "difficult", "disgusting", "distinct", "disturbed", "dizzying", "drab", "dull", "eager", "easy", "elated", "elegant", "embarrassed", "enchanted", "encouraging", "energetic", "enthusiastic", "envious", "evil", "exciting", "expensive", "exuberant", "faithful", "famous", "fancy", "fantastic", "fierce", "filthy", "fine", "foolish", "fragile", "frail", "frantic", "friendly", "frightening", "funny", "gentle", "gifted", "glamorous", "gleaming", "glorious", "good", "gorgeous", "graceful", "green", "grieving", "grumpy", "handsome", "happy", "healthy", "helpful", "helpless", "hilarious", "homeless", "horrible", "hungry", "hurt", "ill", "important", "impossible", "impromptu", "improvised", "inexpensive", "innocent", "inquiring", "itchy", "jealous", "jittery", "joyous", "kind", "knightly", "lazy", "lemony", "light", "lingering", "lively", "lonely", "long", "lovely", "lucky", "magnificent", "modern", "motionless", "muddy", "mushy", "mysterious", "naughty", "niche", "nervous", "nice", "nutty", "obedient", "obnoxious", "odd", "open", "orange", "outrageous", "outstanding", "panicked", "perfect", "pink", "plain", "pleasant", "poised", "poor", "powerless", "precious", "prickling", "proud", "purple", "puzzled", "quaint", "queer", "quizzical", "realistic", "red", "relieved", "repelling", "repulsive", "rich", "scary", "scenic", "selfish", "shiny", "shy", "silly", "sleepy", "smiling", "smoggy", "sore", "sparkly", "splendid", "spotted", "stormy", "strange", "stupid", "successful", "super", "talented", "tame", "tasty", "tender", "tense", "terse", "terrible", "thankful", "thoughtful", "tired", "tough", "troubling", "ugly", "uninterested", "unusual", "upset", "uptight", "varied", "vast", "victorious", "wandering", "weary", "white", "wicked", "wide", "wild", "witty", "worrisome", "wrong", "yellow", "young", "zealous"},
	"adjective_quantitative":  {"a little", "a little bit", "a lot", "abundant", "all", "any", "couple", "double", "each", "either", "empty", "enough", "enough of", "every", "few", "full", "great", "half", "heavily", "heavy", "huge", "hundred", "hundreds", "insufficient", "light", "little", "lots of", "many", "most", "much", "neither", "no", "numerous", "plenty of", "several", "significant", "single", "so few", "some", "sparse", "substantial", "sufficient", "too", "whole"},
	"adjective_proper":        {"Afghan", "African", "Alaskan", "Alpine", "Amazonian", "American", "Antarctic", "Aristotelian", "Asian", "Atlantean", "Atlantic", "Bahamian", "Bahrainean", "Balinese", "Bangladeshi", "Barbadian", "Barcelonian", "Beethovenian", "Belgian", "Beninese", "Bismarckian", "Brazilian", "British", "Buddhist", "Burkinese", "Burmese", "Caesarian", "Californian", "Cambodian", "Canadian", "Chinese", "Christian", "Colombian", "Confucian", "Congolese", "Cormoran", "Costa Rican", "Cypriot", "Danish", "Darwinian", "Diabolical", "Dutch", "Ecuadorian", "Egyptian", "Einsteinian", "Elizabethan", "English", "Finnish", "French", "Freudian", "Gabonese", "Gaussian", "German", "Greek", "Guyanese", "Himalayan", "Hindu", "Hitlerian", "Honduran", "Icelandic", "Indian", "Indonesian", "Intelligent", "Iranian", "Iraqi", "Italian", "Japanese", "Jungian", "Kazakh", "Korean", "kuban", "Kyrgyz", "Laotian", "Lebanese", "Lilliputian", "Lincolnian", "Machiavellian", "Madagascan", "Malagasy", "Marxist", "Mayan", "Mexican", "Middle Eastern", "Monacan", "Mozartian", "Muscovite", "Nepalese", "Newtonian", "Norwegian", "Orwellian", "Pacific", "Parisian", "Peruvian", "Philippine", "Plutonian", "Polish", "Polynesian", "Portuguese", "Putinist", "Roman", "Romanian", "Rooseveltian", "Russian", "Salvadorean", "Sammarinese", "Senegalese", "Shakespearean", "Slovak", "Somali", "South American", "Spanish", "Spanish", "Sri-Lankan", "Sudanese", "Swazi", "Swiss", "Taiwanese", "Thai", "Thatcherite", "Tibetan", "Torontonian", "Turkish", "Turkishish", "Turkmen", "Uzbek", "Victorian", "Viennese", "Vietnamese", "Welsh"},
	"adjective_demonstrative": {"this", "that", "these", "those", "it", "here", "there", "over there"},
	"adjective_possessive":    {"my", "your", "his", "her", "its", "our", "their"},
	"adjective_interrogative": {"what", "whose", "where", "why", "how", "which"},
	"adjective_indefinite":    {"all", "any", "anything", "everyone", "few", "nobody", "one", "some", "someone", "everybody", "anyone", "each", "everything", "many", "none", "several", "somebody"},

	// Pronouns
	"pronoun_personal":      {"I", "we", "you", "he", "she", "it", "they"},
	"pronoun_object":        {"me", "us", "you", "her", "him", "it", "them"},
	"pronoun_possessive":    {"mine", "ours", "yours", "hers", "his", "theirs"},
	"pronoun_reflective":    {"myself", "yourself", "herself", "himself", "itself", "ourselves", "yourselves", "themselves"},
	"pronoun_indefinite":    {"all", "another", "any", "anybody", "anyone", "anything", "both", "each", "either", "everybody", "everyone", "everything", "few", "many", "most", "neither", "nobody", "none", "no one", "nothing", "one", "other", "others", "several", "some", "somebody", "someone", "something", "such"},
	"pronoun_demonstrative": {"this", "that", "these", "those"},
	"pronoun_interrogative": {"who", "whom", "which", "what", "whose", "where", "when", "why", "how"},
	"pronoun_relative":      {"as", "that", "what", "whatever", "which", "whichever", "who", "whoever", "whom", "whomever", "whose"},

	// Connectives
	"connective_time":        {"after a while", "afterwards", "at once", "at this moment", "at this point", "before that", "finally", "first", "here", "in the end", "lastly", "later on", "meanwhile", "next", "next time", "now", "on another occasion", "previously", "since", "soon", "straightaway", "then", "until then", "when", "whenever", "while"},
	"connective_comparitive": {"additionally", "also", "as well", "even", "furthermore", "in addition", "indeed", "let alone", "moreover", "not only", "alternatively", "anyway", "but", "by contrast", "differs from", "elsewhere", "even so", "however", "in contrast", "in fact", "in other respects", "in spite of this", "in that respect", "instead", "nevertheless", "on the contrary", "on the other hand", "rather", "though", "whereas", "yet", "after all", "anyway", "besides", "moreover"},
	"connective_complaint":   {"besides", "e.g.", "for example", "for instance", "i.e.", "in other words", "in that", "that is to say"},
	"connective_listing":     {"firstly", "secondly", "first of all", "finally", "lastly", "for one thing", "for another", "in the first place", "to begin with", "next", "in summation", "to conclude"},
	"connective_casual":      {"accordingly", "all the same", "an effect of", "an outcome of", "an upshot of", "as a consequence of", "as a result of", "because", "caused by", "consequently", "despite this", "even though", "hence", "however", "in that case", "moreover", "nevertheless", "otherwise", "so", "so as", "stemmed from", "still", "then", "therefore", "though", "under the circumstances", "yet"},
	"connective_examplify":   {"accordingly", "as a result", "as exemplified by", "consequently", "for example", "for instance", "for one thing", "including", "provided that", "since", "so", "such as", "then", "therefore", "these include", "through", "unless", "without"},
}
