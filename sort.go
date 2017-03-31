package learn

type SeqSort struct {
	Data []int
}

func (s *SeqSort) BubbleSort() {
	for i := 0; i < len(s.Data); i++ {
		for j := i + 1; j < len(s.Data); j++ {
			if s.Data[i] < s.Data[j] {
				s.Data[i], s.Data[j] = s.Data[j], s.Data[i]
			}
		}
	}
}

func (s *SeqSort) PromBubbleSort() {
	flag := true
	for i := 0; i < len(s.Data) && flag == true; i++ {
		flag = false
		for j := len(s.Data) - 2; j >= i; j-- {
			if s.Data[j] < s.Data[j+1] {
				s.Data[j], s.Data[j+1] = s.Data[j+1], s.Data[j]
				flag = true
			}
		}
	}
}

func (s *SeqSort) SelectSort() {

	for i := 0; i < len(s.Data); i++ {
		max := i
		for j := i + 1; j < len(s.Data); j++ {
			//fmt.Println(i, j, min)
			if s.Data[j] > s.Data[i] {
				max = j
			}
		}
		if i != max {
			s.Data[i], s.Data[max] = s.Data[max], s.Data[i]
		}
	}
}

//插入排序假设前面都已经有序

func (s *SeqSort) InsertSort() {

	for i := 0; i < len(s.Data)-1; i++ {
		if s.Data[i] < s.Data[i+1] {
			//tmp := s.Data[i+1]
			for j := i; j >= 0; j-- {
				if s.Data[j] < s.Data[j+1] {
					s.Data[j], s.Data[j+1] = s.Data[j+1], s.Data[j]
				} else {
					//前面已经排好序无需再次判断
					break
				}
			}
		}
	}
}

//希尔排序关键在于增量的选择
func (s *SeqSort) ShelSort() {
	incre := len(s.Data)
	for incre > 0 {
		incre = incre / 2
		for i := incre; i < len(s.Data); i++ {
			tmp := i - incre
			if s.Data[i] > s.Data[tmp] {
				s.Data[i], s.Data[tmp] = s.Data[tmp], s.Data[i]
			}
		}
	}
}

//快速排序
func (s *SeqSort) QuickSort() {
	Qsort(s, 0, len(s.Data)-1)
}
func Qsort(s *SeqSort, low, high int) {
	var pivot int
	if low < high {
		pivot = Partition(s, low, high)
		Qsort(s, low, pivot-1)
		Qsort(s, pivot+1, high)
	}
}
func Partition(s *SeqSort, low, high int) int {
	var pivotkey int
	pivotkey = s.Data[low]
	for low < high {
		for low < high && s.Data[high] >= pivotkey {
			high = high - 1
		}
		s.Data[low], s.Data[high] = s.Data[high], s.Data[low]
		for low < high && s.Data[low] <= pivotkey {
			low++
		}
		s.Data[low], s.Data[high] = s.Data[high], s.Data[low]
	}
	return low
}
