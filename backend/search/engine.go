package search

import (
	"backend/models"
	"strings"
	"time"
)

type InvertedIndex map[string][]int

type SearchEngine struct {
	Entries []models.LogEntry
	Index   InvertedIndex
}

func NewSearchEngine(entries []models.LogEntry) *SearchEngine {
	engine := &SearchEngine{Entries: entries}
	engine.buildIndex()
	return engine
}

func (s *SearchEngine) buildIndex() {
	s.Index = make(InvertedIndex)
	for i, entry := range s.Entries {
		text := strings.ToLower(entry.Message + " " + entry.MessageRaw + " " + entry.StructuredData +
			" " + entry.Tag + " " + entry.Sender + " " + entry.Groupings + " " +
			entry.Event + " " + entry.Namespace)

		for _, word := range strings.Fields(text) {
			word = strings.Trim(word, ".,:;!?\"'()[]{}")
			s.Index[word] = append(s.Index[word], i)
		}
	}
}

func (s *SearchEngine) Search(query string) ([]models.LogEntry, int, time.Duration) {
	start := time.Now()
	query = strings.ToLower(query)

	ids, found := s.Index[query]
	if !found {
		return nil, 0, time.Since(start)
	}

	results := make([]models.LogEntry, 0, len(ids))
	for _, id := range ids {
		results = append(results, s.Entries[id])
	}

	return results, len(results), time.Since(start)
}
