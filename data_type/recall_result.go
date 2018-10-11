package data_type

import (
  "log"
  "sort"
)

type RecallResult struct {
  Video_id   int64                `json:"video_id"`
  Recall_source []RecallSource    `json:"recall_source"`
}

type By func(p1, p2 *RecallResult) bool

func (by By) Sort(recall_result []RecallResult) {
  sorter := &RecallResultSorter{
    recall_result: recall_result,
    by:      by,
  }
  sort.Sort(sorter)
}

type RecallResultSorter struct {
  recall_result   []RecallResult
  by              func(p1, p2 *RecallResult) bool
}

// Len is part of sort.Interface.
func (s *RecallResultSorter) Len() int {
  return len(s.recall_result)
}

// Swap is part of sort.Interface.
func (s *RecallResultSorter) Swap(i, j int) {
  s.recall_result[i], s.recall_result[j] = s.recall_result[j], s.recall_result[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (s *RecallResultSorter) Less(i, j int) bool {
  return s.by(&s.recall_result[i], &s.recall_result[j])
}

type CompareManager struct {
  P1 []RecallResult
  P2 []RecallResult
}

func(cm *CompareManager) Match() bool {

  if (cm.P1 == nil || cm.P2 == nil) {
    log.Println("either json slice is empty...")
    return false
  }

  if (len(cm.P1) != len(cm.P2)) {
    log.Printf("two file element length is different. P1: [%d], P2: [%d]",
               len(cm.P1), len(cm.P2))
    return false
  }

  // sort by id
  id_sorter := func(p1, p2 *RecallResult) bool {
    return p1.Video_id < p2.Video_id
  }

  By(id_sorter).Sort(cm.P1)
  By(id_sorter).Sort(cm.P2)

  for i := 0 ; i < len(cm.P1); i++ {
    res1 := cm.P1[i]
    res2 := cm.P2[i]
    if (res1.Video_id != res2.Video_id) {
      log.Printf("The [%d]th element is not equal, id1:[%d], id2:[%d]", i + 1, res1.Video_id, res2.Video_id)
      return false
    }

    // sort by recall source name
    SortRecallSource(res1.Recall_source)
    SortRecallSource(res2.Recall_source)

    if (len(res1.Recall_source) != len(res2.Recall_source)) {
      log.Printf("The [%d]th element is not equal, source_size1:[%d], source_size2:[%d]",
                 i + 1, len(res1.Recall_source), len(res2.Recall_source))
      return false
    }

    for j := 0; j < len(res1.Recall_source); j++ {
      source1 := res1.Recall_source[j]
      source2 := res2.Recall_source[j]
      if (source1.Name != source2.Name || source1.Score != source2.Score) {
        log.Printf("The [%d]th element id [%d],  [%d]th source is not equal, name1:[%s], score1:[%f], name2:[%s], score2:[%f]",
                   i + 1, res1.Video_id, j + 1, source1.Name, source1.Score,
                   source2.Name, source2.Score)
        return false
      }
    }
  }
  return true
}
