package data_type

import (
  "sort"
)

type RecallSource struct {
  Name   string     `json:"name"`
  Score  float32    `json:"score"`
}

func SortRecallSource(recall_source []RecallSource) {
  sorter := &RecallSourceSorter{
    recall_source: recall_source,
    by:      SortByName,
  }
  sort.Sort(sorter)
}

type RecallSourceSorter struct {
  recall_source   []RecallSource
  by              func(p1, p2 *RecallSource) bool
}

// Len is part of sort.Interface.
func (s *RecallSourceSorter) Len() int {
  return len(s.recall_source)
}

// Swap is part of sort.Interface.
func (s *RecallSourceSorter) Swap(i, j int) {
  s.recall_source[i], s.recall_source[j] = s.recall_source[j], s.recall_source[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *RecallSourceSorter) Less(i, j int) bool {
  return s.by(&s.recall_source[i], &s.recall_source[j])
}

func SortByName(p1, p2 *RecallSource) bool {
  return p1.Name < p2.Name;
}
