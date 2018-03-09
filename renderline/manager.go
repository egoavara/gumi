package renderline

import (
	"image"
	"sync"
	"golang.org/x/image/draw"
)

type Manager struct {
	// R_
	// 루트 노드를 저장함
	Root *Node
	// R_
	// 이미지를 참조하기 위해 정해짐, 기본적으로 Render()메소드의 결과는 여기에 저장
	completeImage *image.RGBA
	decalRect image.Rectangle
	baseImage *image.RGBA
	decalImage *image.RGBA
	// __
	wgpool *sync.Pool
}
// Node의 Setup과정에 도움을 주는 메서드
// 편의성 이외의 의미는 없음
func (s *Manager ) New(node *Node) *Node {
	temp := &Node{
		_NodeStructure {Manager: s, Parent: node, Childrun: nil},
		_NodeData {},
		_NodeRenderInfomation{},
	}
	if node == nil{
		temp.Allocation = s.completeImage.Rect
		s.Root = temp
	}else {
		temp.Allocation = node.Allocation
		node.Childrun = append(node.Childrun, temp)
	}

	return temp
}
// Setup은 이미지의 크기가 변경될 때마다 반드시 이뤄져야 함.
// 이미지 리사이징이 이뤄진 경우 반드시 주의할 것
func (s *Manager ) Setup()  {
	s.Root.Setup()
}
func (s *Manager ) Render()  {
	s.decalRect = image.ZR
	wg := s.wgpool.Get().(*sync.WaitGroup)
	defer s.wgpool.Put(wg)

	wg.Add(2)
	go func() {
		// 모든 요소들이 캐싱되 있고(즉 변경된 내용이 하나도 없는 경우) 별도의 렌더링 작업이 필요치 않다고 판단되는 경우
		if !AllCached(s.Root){
			s.Root.BaseRender()
		}
		wg.Done()
	}()
	go func() {
		s.Root.DecalRender(&s.decalRect)
		wg.Done()
	}()
	wg.Wait()

}
func (s *Manager) Width() (int) {
	return s.completeImage.Rect.Dx()
}
func (s *Manager) Height() (int) {
	return s.completeImage.Rect.Dy()
}
func (s *Manager) Size() (w, h int) {
	return s.completeImage.Rect.Dx(), s.completeImage.Rect.Dy()
}
func (s *Manager) Rect() (image.Rectangle) {
	return s.completeImage.Rect
}
func (s *Manager) Image() *image.RGBA {
	if s.decalRect == image.ZR {
		return s.baseImage
	}
	draw.Draw(s.completeImage, s.completeImage.Rect, s.baseImage, s.baseImage.Rect.Min, draw.Src)
	draw.Draw(s.completeImage, s.decalRect, s.decalImage, s.decalRect.Min, draw.Over)
	return s.completeImage
}
func NewManager(w, h int) *Manager {
	sz := image.Rect(0,0,w,h)
	return &Manager{
		completeImage: image.NewRGBA(sz),
		baseImage:     image.NewRGBA(sz),
		decalImage:    image.NewRGBA(sz),
		wgpool:&sync.Pool{
			New: func() interface{} {
				return new(sync.WaitGroup)
			},
		},
	}
}
func AllCached(nd *Node) bool {
	if !nd.CacheValid{
		return false
	}
	for _, child := range nd.Childrun{
		if !AllCached(child){
			return false
		}
	}
	return true
}
