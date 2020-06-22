Perfume 구조와 개발 가이드
======================
## 0. 파싱 Format Tree 뷰어
--------------------
```
window := NewWindow(NewSize(32, 80))
body := NewBody(NewSize(32, 80), "MainBody")
stack := NewLayout(StackLayoutType,"MyLayout")
input := NewElement(InputElementType, "MyInput", NewRelativeLocation(5, 10))

_ = body.AddChild(stack)
_ = stack.AddChild(input)
_ = window.Add(body)

renderer := NewRenderer(window)

renderer.PrintStruct(ElementsPrintDepth, map[PrintLineForm]*Parseable{
	WindowLine:   NewParseable("Window || (%Size%) || (%ChildrenLen%) ||\n\n"),
	FormalsLine:  NewParseable("-- (%Name%) Formal --\n"),
	LayoutsLine:  NewParseable("\t└--(%Type%)layout %Name%\n"),
	ElementsLine: NewParseable("\t\t└--(%Type%)element LOC:%RelLocation%\n"),
})
```
위와 같은 코드로 Renderer를 만든 후, PrintStruct함수로 Window와 그 하위 객체들에 대한 그림을 그리나, %PROPERTY%로 묶은 문자열로, 자율적으로 프린팅Printing 방법을 고칠 수 있습니다. 더 자세한 부분은 parser.go의 PrintPropertyType를 봐주세요.

## 1.구조
--------------

퍼퓸은 총 4개의 엘리먼트(Element)들로 이루어져 있습니다.  
상위 엘리먼트부터 나열하자면 Window, Formal,Layout, Element가 있습니다.  
## 특징
------------

1. ### Window 엘리먼트는 형식이 1개입니다.  
---
절대로 변경될 수 없으며, 빌더(Builder) 역할을 합니다.  
Window에는 Children이 있는데, 그것들은 모두 IFormal입니다.  

2. ### Formal 엘리먼트는 Window보다 작거나 같아야합니다.  
---
Formal 엘리먼트는 ILayout들만을 Children으로 가집니다.  

3. ### Layout 엘리먼트는 2가지 입니다.  
----

첫째, 스택 레이아웃  
이것은 자식의 RelativeLocation 속성을 무력화시킵니다.  
자신의 Orientation 속성에 따라서 가로, 세로 스택의 형태만 가집니다.  

둘째, 프리 레이아웃
이것은 자식의 RelativeLocation 속성을 고려해 자식의 위치를 정합니다. 그 외 특별한 속성은 존재 하지 않습니다.  

셋째, 부모 엘리먼트로써 IFormal을 가지며, 자식 엘리먼트를 IElement의 형태로 가집니다.

4. ### Element 엘리먼트
----

부모 엘리먼트로 ILayout을 가지며 다양한 속성이 적용 될 수 있고, 가장 많이 커스텀 엘리먼트가 발생될 수 있습니다.

5. ### 인터페이스 IFormal, ILayout, IElement
-----
이것 외에 iFormalElement, iLayoutElement, iElement등이 존재하는데, 이것은 구조체 FormalElement, LayoutElement, Element의 인터페이스 들입니다. 그리고 이 구조체들은 FreeLayout, Input 등의 엘리먼트들의 필수요소가 됩니다.   
> 즉, Head 구조체에는 FormalElement가 있는데, FormalElement는 iFormalElement를 따르니까, iFormalELement 형식의 변수에 Head의 FormalElement(Head.FormalElement입니다. Head가 아닙니다!)가 들어갈 수 있습니다. 그런데 이렇게 하면 각 커스텀 엘리먼트들이나 Input, Text와 같은 엘리먼트들에 특성을 부여하기 매우 어렵습니다. 그래서 ILayout과 IFormal, IElement와 같이 iFormal .. iLayout ... iElement 의 필수 함수들을 구현한것을 포함하는 새로운 인터페이스를 만들었습니다. 이렇게 한다면, Input 과 같은 구조체에 따로 InputStyles와 같은 함수 등을 추가 하여도, IElement형식의 변수에 저장이 가능합니다.  
>>결론적으로, 소문자 i 로 시작하는 인터페이스들은 모든 엘리먼트들이 기본으로 갖춰야 할 것들을 선언하며, LayoutElement, Element, FormalElement들은 이것들을 구현하고 있습니다. 따라서 우리는, 커스텀 엘리먼트를 생성할때 원하는 엘리먼트에 따라 코드 구현이 가능합니다.
```
type Input struct {
	kind InputType
    //불러온 Element구조체가 이미 iElement인터페이스의 필수 함수들을 모두 구현한 상태입니다.
	Element
}
```
>>위 코드 처럼 Input 구조체를 구현 하며 Element 규칙을 따를 수 있습니다.

type.go에서  ~ElementType의 상수들의 구조체를 매개변수로 받고싶을땐, IFormal, ILayout, IElement 인터페이스를 사용하여 받을 수 있습니다.
마지막으로 쉽게 말해 구조체와 인터페이스들은 아래와 같이 대응됩니다.
```
FormalElement -> iFormalElement
LayoutElement -> iLayoutElement
Element -> iElement

//IFormal을 구현한 것은 필연적으로 iFormalElement를 구현합니다
IFormal -> iFormalElement

ex)
Footer == IFormal                      // true
Footer.FormalElement == iFormalElement // true
```

## 3. 추후 개발 가이드( 변할 수 있습니다! )
----------------
1. 모든 엘리먼트(Formal,Layout 등)은 이름을 가져야 합니다  

2. 자식으로 속해있을때 그들 중에 중복된 이름을 가져선 안됩니다.  

3. 읽을땐 무조건 부모 -> 자식 순으로 읽습니다.

4. 임의로 함수를 거치지 않고 부모, 자식을 사용자(개발자)가 바꾸어선 안됩니다.