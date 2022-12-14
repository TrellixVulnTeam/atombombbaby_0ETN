GOOF----LE-4-2.0?      ] ; 4   h      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  sxml?	g  	transform?	 ?		g  filenameS?	
f  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?	g  exportsS?	g  SRV:send-reply?	g  foldts?	g  
post-order?	g  pre-post-order?	g  replace-range?	 ?	g  set-current-module?	 ?	 ?	g  make-syntax-transformer?	 ?	 ?	g  let*-values?	g  macro?	g  $sc-dispatch?	 ?	 ?	g  _?	g  any?	??	 g  syntax->datum?	!  ?	"  ?	#g  datum->syntax?	$# ?	%# ?	&g  begin?	'g  let?	(g  call-with-values?	)g  lambda?	*g  syntax-violation?	+* ?	,* ?	-f  -source expression failed to match any pattern?	.g  
procedure??	/g  display?	0g  assq?	1g  	*default*?	2g  *text*?	3g  
*preorder*?	4g  *macro*?	5g  append?	6g  map?	7g  error?	8f  Unknown binding for ?	9f   and no default?	:g  reverse?C 5      h?  ?   ]4	
5 4 >  "  G   4"%&'()        hX   7  ]
L ?L?? ?$  1 ?(   ?   C    C    C /      g  vars
		W g  initializer		W g  cont			W  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
	V	??		X	??		Y	??	
	Y	??		X	??		W		??		[	??		Z	??		]	??		Z	??		^	??	 	^	??	%	^	??	0	`	??	G	\	?? 		W	   C      h(   ?   - 1 3  (  ?C O  ?@     ?       g  bindings
			# g  body			#  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
	S
??		T	??		T	??		T	??	!	b	??	#	U	?? 			#
   C   h   ?   ]	4 5L 4?6?       g  args
		 g  v			  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?		S
?? 		   C,-       h(   ?   ]	4 5$   O @ 6 ?       g  y
		' g  tmp		'  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
	S
?? 		'   C5R./  h?   ?  ] (  C ?$  ? ?(   ? "??? ?&   ? "??? ??$   ?4L  ?5 "???4 ?5$  4 ?>   "  G   ? "???4 ?>  "  G   ? "??c ? "??V    ?      g  	fragments
	 ? g  result	 ?  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
	n	??		o	??		q	??		o	??		r	??		o	??		r	%??	&	r	??	*	s	??	.	o	??	1	s	&??	:	s	 ??	=	t	??	>	t	??	B	o	??	E	u	??	F	u	??	K	u	$??	O	u	??	W	u	??	X	v	??	]	v	??	_	v	??	c	o	??	d	w	??	g	w		??	l	w	??	w	x	?? ?	x	?? ?	z	?? ?	z	?? ?	z	?? ?	{	?? ?	{	?? ?	q	#?? ?	q	?? %	 ?	  g  nameg  loop? C    h    ?  -  1  3 O Q  6   ?      g  	fragments
			 g  loop		  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
	d
??		n	?? 			


  g  nameg  SRV:send-reply?g  documentationf Output the @var{fragments} to the current output port.

The fragments are a list of strings, characters, numbers, thunks,
@code{#f}, @code{#t} -- and other fragments. The function traverses the
tree depth-first, writes out strings and characters, executes thunks,
and ignores @code{#f} and @code{'()}. The function returns @code{#t} if
anything was written at all; otherwise the result is @code{#f} If
@code{#t} occurs among the fragments, it is not written out but causes
the result of @code{SRV:send-reply} to be @code{#t}.? CR012.034567892   h?   ?  ] (  C ?$  ? ??$  ? ?4L 5$  "  L$  _??$  F??&  ?? @??&  4?? ? "?????4 ?4??L 55@?4L ?5@	6L 6L$  L
 6
	6     ?      g  tree
	 ? g  trigger	 ? g  t		'	: g  binding		: ?  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
 ?	??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	
??	 ?	??	' ?	??	: ?	
??	B ?	??	E ?	??	F ?	??	J ?	??	L ?	??	O ?	 ??	T ?	??	W ?	??	\ ?	??	^ ?	??	a ?	??	f ?	??	g ?	??	j ?	??	o ?	??	u ?	??	x ?	??	| ?	?? ? ?	"?? ? ?	-?? ? ?	5?? ? ?	-?? ? ?	?? ? ?	?? ? ?	?? ? ?	-?? ? ?	7?? ? ?	-?? ? ?	?? ? ?	?? ? ?	6?? ? ?	?? ? ?	$?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	4?? ? ?	?? 6	 ?  g  nameg  loop? C hp   ?  ]"4545$  "  $  4?5$  ?"  ??"  O Q  6     ?      g  tree
		k g  bindings		k g  default-binding			k g  t			) g  text-binding		)	k g  text-handler		M	k g  loop		Z	k  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
 ?
??	 ?	??	 ?	 ??	 ?	??	 ?	??	 ?	??	 ?	!??	 ?	??	 ?	??	) ?	??	1 ?	??	2 ?	??	7 ?	??	9 ?	??	= ?	??	@ ?	??	G ?	$??	M ?	??	Z ?	?? 		k	  g  nameg  pre-post-order? CRiR       h`   ?  ](  C?$  C"  +(  
64 ?5?"???4 5?"???6   ?      g  fdown
		] g  fup		] g  fhere			] g  seed			] g  tree			] g  kid-seed			A g  kids			A  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
 ?
??	 ?	??	 ?		??	 ?	??	 ?	??	 ?	??	& ?	
??	' ?	??	4 ?	1??	6 ?	??	9 ?	??	A ?	
??	A ?	??	B ?	??	M ?	2??	U ?	??	] ?	?? 		]	  g  nameg  foldts? CR:5    hP  ?  ]+ (  45D ?$  ?4L 5$   ?4455 "????$  K??4L$  ?"  >  G  ?$  
??"  ? "??g ?? "??S4L5$  4 ?5 "??-?$  c??4L$  ?"  >  G  ?$  "  ?$  $  
??"  ?"   "??? ? "????      g  forest
	P g  keep?	P g  
new-forest		P g  node		P g  t		% ? g  node?		T ? g  new-kids		p ? g  keep?		p ? g  t	 ?P g  node?	 ?? g  new-kids	 ?? g  keep?	 ??  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
 ?	??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	 ?	
??		??	% 	??	0	??	2	??	5	 ??	?	??	I	??	L	??	P 	??	S		$??	T		??	T	??	W	??	_	??	b	#??	j	7??	o	??	z	?? ?	?? ?	(?? ?	"?? ?	?? ?	?? ?	?? ?	)?? ?	?? ?	?? ?	?? ?	?? ?	.?? ?	?? ?	?? ?	?? ?	?? ?	$?? ?	?? ?	?? ?	?? ?	?? ?	#?? ?	7?? ?	??	??	??	%??	??  	??# 	,??& 	&??/	???	??B	??P	?? =	P	  g  nameg  loop? C   h0   E  ]O  Q 4>  G C    =      g  beg-pred
		, g  end-pred		, g  forest			, g  loop			, g  
new-forest		"	, g  keep?		"	,  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?
 ?
??	%	$??	%	4??	!%	$??	%%	?? 		,	  g  nameg  replace-range? CRC      ?       g  m
		(  g  filenamef  V/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/transform.scm?		H
??
?	d
??? ?
??? ?
??? ?
??? ?
?? 	?
   C6 