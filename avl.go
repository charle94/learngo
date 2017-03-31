package learn

/*
type avlTree struct {
	Tree
	Data       int
	Leftchild  *avlTree
	Rightchild *avlTree
	bf         int
}

const (
	LH = 1
	EH = 0
	RH = -1
)

func (avl *avlTree) R_Roate() {
	tmp := avl.Leftchild
	avl.Leftchild = tmp.Rightchild
	tmp.Rightchild = avl
	avl = tmp
}
func (avl *avlTree) L_Roate() {
	tmp := avl.Rightchild
	avl.Rightchild = tmp.Leftchild
	tmp.Leftchild = avl
	avl = tmp
}*/

/*func (avl *avlTree) LeftBlance() {
	var Lr *avlTree
	L := avl.Leftchild
	switch bf := L.bf; bf {
		case LH:
			avl.bf=L.bf=EH
			avl.R_Roate()
		case RH:
			Lr:=L.Rightchild
			switch Lr.bf{
				case LH:
					avl.bf=RH
					L.bf=EH
				case EH:
					avl.bf=L.bf=EH
				case RH:
					avl.bf=EH
					L.bf=LH

			}
			Lr.bf=EH
			avl.Rightchild.L_Roate()
			avl.R_Roate()
	}

}
func (avl *avlTree) RightBlance() {
	var Rl * avlTree
	R:=avl.Rightchild
	switch R.bf{
	case RH :
		avl.bf=R.bf=EH
		avl.L_Roate()
	case LH:
		Rl=R.Leftchild
		switch Rl.bf{
		case EH:
			avl.bf=R.bf=EH
		case LH:
			avl.bf=EH
			R.bf=RH
		case RH:
			avl.bf=LH
			R.bf=EH
		}
		Rl.bf=EH
		avl.Leftchild.R_Roate()
		avl.L_Roate()
	}
}*/
