GOOF----LE-4-2.0_+      ] - 4 hÌ      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  sxml¤	g  fold¤	 ¤		g  filenameS¤	
f  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm¤	g  importsS¤	g  srfi¤	g  srfi-1¤	 ¤	 ¤	 ¤	g  exportsS¤	g  foldt¤	g  foldts¤	g  foldts*¤	g  fold-values¤	g  foldts*-values¤	g  fold-layout¤	 ¤	g  set-current-module¤	 ¤	 ¤	g  atom?¤	g  map¤	g  apply¤	g  append¤	 g  assq¤	!g  assq-ref¤	"g  error¤	#f  no binding available¤	$g  @¤	%g  macro¤	&g  pre¤	'g  reverse¤	(g  bindings¤	)g  
pre-layout¤	*g  post¤	+g  	*default*¤	,g  *text*¤C 5       hh'  Õ   ]4	
5 4 >  "  G     h   µ   ] C ­       g  x
		  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	&
		'			'	 		  g  nameg  atom? CR     h      ]LL  6           g  kid
		  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	0			1	 		   C       h(   t  ]45$  6 4 O 56 l      g  fup
		' g  fhere		' g  tree			'  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	)
		.			.			/			0		'	0	 		'	  g  nameg  foldtg  documentationf  jThe standard multithreaded tree fold.

@var{fup} is of type [a] -> a. @var{fhere} is of type object -> a.
 CR      h   ±   ]LLL  6©       g  kid
		 g  kseed		  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	:			;	 			   C    h8   ®  ]45$  64 O 4 556¦      g  fdown
		8 g  fup		8 g  fhere			8 g  seed			8 g  tree			8  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	4
		7			7			8			:		(	<		4	:		8	9	 			8	  g  nameg  foldtsg  documentationf  nThe single-threaded tree fold originally defined in SSAX.
@xref{sxml ssax,,(sxml ssax)}, for more information. CR    h   ±   ]LLL  6©       g  kid
		 g  kseed		  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	J			K	 			   C    hH     ]45$  64 >  G 4 O 56         g  fdown
		E g  fup		E g  fhere			E g  seed			E g  tree			E g  kseed		"	E g  tree		"	E  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	@
		D			D			E			G		%	F		.	J		E	I	
 			E	  g  nameg  foldts*g  documentationf  ²A variant of @ref{sxml fold foldts,,foldts} that allows pre-order
tree rewrites. Originally defined in Andy Wingo's 2007 paper,
@emph{Applications of fold to XML transformation}. CR     h8   Ò  - 1 3 (  E4 >  G @     Ê      g  proc
			3 g  list			3 g  seeds				3 g  seeds		#	3  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	Q
		U			V			X			X	!	"	X		&	W		/	Z	"	3	Z	
 
			3	
	  g  nameg  fold-valuesg  documentationf  A variant of @ref{SRFI-1 Fold and Map, fold} that allows multi-valued
seeds. Note that the order of the arguments differs from that of
@code{fold}. CR      h   ³   - 1 3 LLL  @ «       g  tree
			 g  seeds			  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	h			i	 			
   C        h`   E  - 1 3 45$  @4 >  G4 O >  G45@=      g  fdown
			` g  fup			` g  fhere				` g  tree				` g  seeds				` g  tree		+	` g  kseeds		+	` g  kseeds		L	`  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	\
	
	`			`			a			c		.	b		3	g		O	e	
	V	m		`	m	 			`	
	  g  nameg  foldts*-valuesg  documentationf  ¯A variant of @ref{sxml fold foldts*,,foldts*} that allows
multi-valued seeds. Originally defined in Andy Wingo's 2007 paper,
@emph{Applications of fold to XML transformation}. CR   h    ñ   ]4 5$  CC      é       g  alist
		 g  key		 g  default			 g  t			  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	o
		p				p	 			  g  nameg  assq-ref C!R"#   h   ¼   -  1  3  6       ´       g  args
			  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 ¹		 º		 º	 			


  g  nameg  err C h   Æ   -  1  3 LL  @      ¾       g  args
			  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 ½		 ¾	)	 ¾	 			


  g  nameg  cont-with-tag C!$%&'  h   ê   ]45D   â       g  params
		 g  layout		 g  
old-layout			 g  kids			  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 Õ		 Ö		 Ö	 			   C()h   ¿   ]C   ·       g  tag
		 g  params		 g  layout			  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 Ü	 			   C*+  h`  ¦  ]>"  O 	 Q 	$  45"  4L  5
 (  
	D"   
	D $  / &   4 5 
	D"ÿÿ·"ÿÿ³4 5$  45$  !4 ? "ÿÿ45$  4 ?D4544	
5 544L55"ÿþ¿4L5"ÿþ§       g  tree
	_ g  bindings	_ g  pcont		_ g  params		_ g  layout		_ g  ret		_ g  new-bindings		 ¡ g  
new-layout		 ¡ g  cont		 ¡ g  cont-with-tag			 ¡ g  bindings	
	6 ¡ g  style-params		6 ¡ g  params	   g  tag-bindings	 ¬_ g  t	 ¾G g  t	 ñG  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 »		 ¼		 ½		 ¿		 À		, Â		3 Â	/	4 Â	:	6 Â		6 ½		= Ä		A Ã		B Æ		K Æ	&	N Æ	L	P Å	
	W Í		` Í	-	c Í	S	e Ì	
	e Ã		h Ç		j Ç		n Ã		q Ç	(	u Ç	5	y Ç	
	z È		 È	&  È	  È	  È	
  Ê	  Ê	B  É	 ¢ á	 © á	+ ¬ á	 ¬ á	 ´ â	 µ ä	 » ä	 ¾ ä	 ¾ â	 Ç Ñ	 ç Ñ	 è å	 î å	 ñ å	 ñ â	 ú Ô	 ×	 Ô	 Ú	 Ú	 Ú	 Û	 Û	 Û	! Û	) Û	* ß	0 ß	1 à	7 à	$; à	= ß	G Ù	K Ï	Q Ï	1U Ï	_ Ï	 I	_	  g  nameg  fdown C'  h8   Ì  , 3 4	4
5>  G D      Ä      g  tree
		2 g  bindings		2 g  cont			2 g  params			2 g  layout			2 g  ret			2 g  	kbindings			2 g  kcont			2 g  kparams			2 g  klayout				2 g  kret	
		2 g  klayout			2 g  kret			2  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 ç		 ë	
	 ë	(	 ë	
	 é		0 í	-	2 í	 		2	  g  nameg  fup C!,    h0   z  ]44L 5 >  G Dr      g  tree
		0 g  bindings		0 g  cont			0 g  params			0 g  layout			0 g  ret			0 g  tlayout			0 g  tret			0  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
 î		 ñ	
	 ñ		
 ñ		 ñ		 ñ	
	 ï		. ó	-	0 ó	 
		0	  g  nameg  fhere C   hP   ç  ]=O Q 4O   >	  G 
	
D  ß      g  tree
		N g  bindings		N g  params			N g  layout			N g  
stylesheet			N g  err			N g  fdown			N g  bindings		:	N g  cont		:	N g  params			:	N g  layout	
	:	N g  ret		:	N  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm
	s
	 ö		1 ÷	7	4 ÷	C	9 ö		= ô		J ù		N ù	 			N	  g  nameg  fold-layoutg  documentationf 	A traversal combinator in the spirit of SSAX's @ref{sxml transform
pre-post-order,,pre-post-order}.

@code{fold-layout} was originally presented in Andy Wingo's 2007 paper,
@emph{Applications of fold to XML transformation}.

@example
bindings := (<binding>...)
binding  := (<tag> <bandler-pair>...)
          | (*default* . <post-handler>)
          | (*text* . <text-handler>)
tag      := <symbol>
handler-pair := (pre-layout . <pre-layout-handler>)
          | (post . <post-handler>)
          | (bindings . <bindings>)
          | (pre . <pre-handler>)
          | (macro . <macro-handler>)
@end example

@table @var
@item pre-layout-handler
A function of three arguments:

@table @var
@item kids
the kids of the current node, before traversal
@item params
the params of the current node
@item layout
the layout coming into this node
@end table

@var{pre-layout-handler} is expected to use this information to return a
layout to pass to the kids. The default implementation returns the
layout given in the arguments.

@item post-handler
A function of five arguments:
@table @var
@item tag
the current tag being processed
@item params
the params of the current node
@item layout
the layout coming into the current node, before any kids were processed
@item klayout
the layout after processing all of the children
@item kids
the already-processed child nodes
@end table

@var{post-handler} should return two values, the layout to pass to the
next node and the final tree.

@item text-handler
@var{text-handler} is a function of three arguments:
@table @var
@item text
the string
@item params
the current params
@item layout
the current layout
@end table

@var{text-handler} should return two values, the layout to pass to the
next node and the value to which the string should transform.
@end table
 CRC    Í       g  m
		,  g  filenamef  Q/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/sxml/fold.scm		
 ø	&
o	)
A	4
		@
­	Q
P	\
t	o
'b	s
 
	'd
   C6 