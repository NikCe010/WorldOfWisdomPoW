package quotes_service

import "math/rand"

// Service - represent service with storage for working with quotes.
//
// It's a cache with preset quotes
type Service map[int]string

// New ...
func New() Service {
	var storage Service = make(map[int]string)

	storage.setDefaultQuotes()

	return storage
}

// GetRandomQuote - method for getting random quote
func (s Service) GetRandomQuote() string {
	return s[rand.Intn(len(s)-1)]
}

func (s Service) setDefaultQuotes() {
	s[0] = "We must believe that we are gifted for something, and that this thing, at whatever cost, must be attained."
	s[1] = "The older I get, the greater power I seem to have to help the world; I am like a snowball - the further I am rolled the more I gain."
	s[2] = "Knowledge is love and light and vision."
	s[3] = "Wisdom is knowing what to do next. Skill is knowing how to do it. Virtue is doing it."
	s[4] = "When one door of happiness closes, another one opens; but we look so long at the closed door that we do not see the one which has opened for us."
	s[5] = "Poor eyes limit your sight; poor vision limits your deeds."
	s[6] = "I do not pray for success. I ask for faithfulness."
	s[7] = "I used to ask God to help me. Then I asked if I might help him."
	s[8] = "The wise person doesn't give the right answers, but poses the right questions."
	s[9] = "What happens is not as important as how you react to what happens."
	s[10] = "Be wise like serpents and harmless like doves."
	s[11] = "The first problem for all of us, men and women, is not to learn, but to unlearn."
	s[12] = "If you have knowledge, let others light their candles in it."
	s[13] = "When the going gets rough - turn to wonder."
	s[14] = "A bird doesn't sing because it has an answer, it sings because it has a song."
	s[15] = "We are not what we know but what we are willing to learn."
	s[16] = "Good people are good because they've come to wisdom through failure."
	s[17] = "Your word is a lamp for my feet, a light for my path."
	s[18] = "By three methods we may learn wisdom: First, by reflection, which is noblest; Second, by imitation, which is easiest; and third by experience, which is the bitterest."
	s[19] = "The reason people find it so hard to be happy is that they always see the past better than it was, the present worse than it is, and the future less resolved than it will be."
}
