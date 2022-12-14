GOOF----LE-4-2.0?      ] [ 4   hf      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  scripts?	g  use2dot?	 ?		g  filenameS?	
f  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?	g  importsS?	g  srfi?	g  srfi-13?	 ?	g  selectS?	g  string-join?	 ?	 ?	g  frisk?	 ?	g  make-frisker?	g  	edge-type?	g  edge-up?	g  	edge-down?	 ?	 ?	 ?	g  exportsS?	 ?	g  	autoloadsS?	g  ice-9?	 g  getopt-long?	!  ?	"  ?	#!" ?	$g  set-current-module?	%$ ?	&$ ?	'f  1Print a module's dependencies in graphviz format.?	(g  %summary?	)g  
guile-user?	*) ?	+g  *default-module*?	,g  format?	-f  ~S?	.g  q?	/g  map?	0f  ~A=~A?	1g  vv?	2f  digraph use2dot {
?	3g  for-each?	4f    ~A;
?	5g  label?	6f  Guile Module Dependencies?	7g  ratio?	8g  fill?	978??	:9 ?	;g  >>header?	<f    "~A" -> "~A"?	=g  autoload?	>g  style?	?g  dotted?	@>???	Ag  fontsize?	BA	??	C@B ?	Dg  computed?	Eg  bold?	F>E??	GF ?	Hf   [~A]?	If  ,?	Jf  ;
?	Kg  >>body?	Lf  }?	Mg  >>footer?	Ng  >>?	Of  use2dot?	Pg  default-module?	Qg  single-char?	RQm ?	Sg  value?	TS ?	UPRT ?	VU ?	Wg  
option-ref?	Xg  reverse?	Yg  edges?	Zg  main?C 5     h  ?   ]4	
#5	 4& >  "  G   '(R*+R,-      h   ?   ] 6     ?       g  s
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	;
??		<	??		<	?? 		  g  nameg  q? C.R/,0   h   ?   ] ? ?6 ?       g  pair
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	?	??		@	??	
	@	??		@	'??		@		?? 		   C        h   ?   ] 6      ?       g  pairs
		
  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	>
??	
	?	?? 		
  g  nameg  vv? C1R,23,4  h   ?   ] 6     ?       g  s
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	E	??		E	#??		E	?? 		   C15.6: 
h0   ?   ] 4>  "  G  445?	?56   ?       g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	C
??		D	??		D	??		D	??		F	??		F	??		F	??	#	F	??	%	F	??	&	F	??	+	F	??	-	E	?? 		-
  g  nameg  >>header? C;R3,<=CDGH1IJ    h?   _  ]	44 54 5>  "  G  4 5?$  "  ?$  	"  $  %4
4455>  "  G  "   6   W      g  edge
	 ? g  key	*	N g  t	N	~  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	O	??		P	??		P	??			P	%??		P	6??		P	??	$	Q	??	*	Q	??	7	R	??	D	Q	??	F	S	??	N	Q	??	W	V	??	\	V	??	]	V	$??	`	V	1??	h	V	:??	j	V	$??	o	V	?? ?	W	?? ?	W	?? 	 ?   C      h   ?   ] 6      ?       g  edges
		
  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	M
??	
	N	?? 		
  g  nameg  >>body? CKR,L     h   ?   ] 6       ?       g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	Z
??		[	??			[	?? 			
  g  nameg  >>footer? CMR;KM     h(   ?   ]4>   "  G  4 >  "  G  6 ?       g  edges
		(  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	]
??		^	??		_	??	(	`	?? 		(  g  nameg  >>? CNR OVWP+NXY hP   ?  -  1  3 4 ?5454?5454	445
556?      g  args
			P g  parsed-args		P g  =m		"	P g  scan		.	P g  files		9	P  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?
	b
??	
	c	??		c	)??		c	#??		d	#??		c	??		c	??		f	??		f	%??	"	f	??	"	c	??	%	g	??	)	g	??	.	g	??	.	c	??	1	h	??	6	h	(??	7	h	,??	9	h	??	9	c	??	>	i	??	A	i	??	B	i	??	J	i	??	L	i	??	N	i	??	P	i	?? 			P


  g  nameg  use2dot? CRiZRC     ?       g  m
		0  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/use2dot.scm?		0
??	2	7	??	5	7
??	7	9	??	:	9
??	;
???	>
???	C
???	M
???	Z
??	?	]
???	b
??	k
?? 	
   C6 