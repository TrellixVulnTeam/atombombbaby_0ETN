GOOF----LE-4-2.00      ] V 4        h4      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  ice-9?	g  
gap-buffer?	 ?		g  filenameS?	
f  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?	g  exportsS?	g  gb??	g  make-gap-buffer?	g  gb-point?	g  gb-point-min?	g  gb-point-max?	g  gb-insert-string!?	g  gb-insert-char!?	g  gb-delete-char!?	g  	gb-erase!?	g  gb-goto-char?	g  
gb->string?	g  
gb-filter!?	g  	gb->lines?	g  gb-filter-lines!?	g  make-gap-buffer-port?	 ?	g  	autoloadsS?	g  srfi?	g  srfi-13?	 ?	 g  string-join?	!  ?	"! ?	#g  set-current-module?	$# ?	%# ?	&g  make-record-type?	'g  s?	(g  all-sz?	)g  gap-ofs?	*g  aft-ofs?	+'()* ?	,g  record-predicate?	-g  record-accessor?	.g  s:?	/g  all-sz:?	0g  gap-ofs:?	1g  aft-ofs:?	2g  record-modifier?	3g  s!?	4g  all-sz!?	5g  gap-ofs!?	6g  aft-ofs!?	7g  default-initial-allocation?	8g  default-chunk-size?	9g  default-realloc-threshold?	:g  round-up?	;g  record-constructor?	<g  new?	=g  make-string?	>g  substring-move!?	?g  realloc?	@g  port??	Ag  eof-object??	Bg  list->string?	Cg  reverse?	Dg  	read-char?	Eg  string??	Fg  string-length?	Gg  error?	Hf  bad init type?	Ig  insert-prep?	Jg  string-set!?	Kg  max?	Lg  min?	Mg  	point++n!?	Ng  	point+-n!?	Og  string-append?	Pg  	substring?	Qg  string-index?	Rf  not a gap-buffer:?	Sg  make-soft-port?	Tg  
string-ref?	Uf  rw?C 5       h?&  ?  ]4	
"5 4% >  "  G   4&i+5R4,ii5R4-ii'5.R4-ii(5/R4-ii)50R4-ii*51R42ii'53R42ii(54R42ii)55R42ii*56R ?7R ?8R	 9R8 h   ?   ] ???C    ?       g  n
		  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??		 ?	??	
 ?	??	 ?	?? 		  g  nameg  round-up? C:R4;ii5<R./01=>346 
    h?   	  ]:4 54 5?4 54 545?4

>  "  G  4>  "  G  4 >  "  G  4 >  "  G  	 6       g  gb
	 ? g  inc	 ? g  old-s			 ? g  all-sz		 ? g  new-sz		 ? g  gap-ofs		" ? g  aft-ofs		+ ? g  new-s		4 ? g  new-aft-ofs		; ?  	g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	" ?	??	% ?	??	+ ?	??	. ?	??	4 ?	??	; ?	??	; ?	??	> ?	??	V ?	??	p ?	?? ? ?	?? ? ?	?? 	 ?	  g  nameg  realloc? C?R<3=7456@ABC:>.DEFGH  h  ?  -  1  3 45  (  X445>  "  G  4>  "  G  4
>  "  G  4>  "  G  " ? ?45$  ?"  ?4	5$  ?4
45545445>  "  G  4>  "  G  4
45
>  "  G  4>  "  G  4>  "  9G  "  245??"??I45
"??6"  ?45$  ?4545445>  "  G  4>  "  G  4
45
>  "  G  4>  "  G  4>  "  G  "  4>  "  G  C       ?      g  init
		 g  gb	 g  v		q? g  c	 ?9 g  acc	 ?9 g  len	 ?9 g  string	 ?  g  alloc	 ? g  len	b? g  alloc	k?  
g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	
 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	' ?	??	0 ?	??	D ?	??	W ?	??	q ?	??	q ?	??	t ?	??	~ ?	?? ? ?	?? ? ?	!?? ? ?	?? ? ?	'?? ? ?	5?? ? ?	'?? ? ?	!?? ? ?	+?? ? ?	?? ? ?	 ?? ? ?	'?? ? ?	 ?? ? ?	 ?? ? ?	 ?? ? ?	>?? ? ?	 ?? ? ?	 ?? ?	 ??! ?	0??, ?	.??/ ?	;??9 ?	!??9 ?	??: ?	0??A ?	6??L ?	??Q ?	??[ ?	??\ ?	!??b ?	??e ?	+??k ?	??n ?	 ??s ?	'??~ ?	 ??? ?	 ??? ?	 ??? ?	>??? ?	 ??? ?	 ??? ?	 ??? ?	??? ?	%??? ?	?? ;		


  g  nameg  make-gap-buffer? CR0  h   ?   ]4 5?C     ?       g  gb
		  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??	
 ?	?? 		  g  nameg  gb-point? CRh   ?   ]C    ?       g  gb
		  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
?? 		  g  nameg  gb-point-min? CR/10      h    ?   ]4 54 54 5???C     ?       g  gb
		  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?		??	
 ?	??	 ?	'??	 ?	??	 ?	??	 ?	?? 		  g  nameg  gb-point-max? CR019?:    hP   ?  ]4 54 5???$  4 4
?5>  "  G  "   C       x      g  gb
		I g  len		I g  gap-ofs				I g  aft-ofs			I g  slack			I  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	# ?		??	' ?	??	( ?		??	- ?	??	0 ?	??	5 ?	??	: ?		?? 		I	  g  nameg  insert-prep? CIRFI>.5       h@   E  ]454 54
4 5>  "  G   ?6 =      g  gb
		? g  string		? g  len				? g  gap-ofs			?  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	 ?	"??	, ?	??	= ?	??	? ?	?? 		?	  g  nameg  gb-insert-string!? CRIJ.5    h0     ]
4 544 5>  "  G   ?6      g  gb
		0 g  char		0 g  gap-ofs		
	0  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??	
 ?	??	 ?	??	 ?	??	 ?	??	. ?	??	0 ?	?? 			0	  g  nameg  gb-insert-char!? CR5K06L/1     hP   d  ]
?$   4
4 5?56
?$   44 54 5?56
?$  CC\      g  gb
		P g  count		P  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?		??	
 ?	??	 ?	??	 ?	 ??	 ?	??	 ?	??	  ?		??	$ ?		??	( ?	??	- ?	??	0 ?	??	7 ?	+??	@ ?	(??	B ?	??	D ?		??	H ?		??	L ?	?? 		P	  g  nameg  gb-delete-char!? CR56/       h(   ?   ]4 
>  "  G   4 56      ?       g  gb
		"  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??	 ?	??	" ?	?? 		"  g  nameg  	gb-erase!? CR>56      hH   B  ]4?>  "  G  4 ?>  "  G   ?6       :      g  gb
		A g  n		A g  s			A g  gap-ofs			A g  aft-ofs			A  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??	 ?	??	 ?	??	  ?	??	) ?	??	. ?	??	? ?	??	A ?	?? 
		A	  g  nameg  	point++n!? CMR>56 hH   K  ]4??>  "  G  4 ?>  "  G   ?6    C      g  gb
		D g  n		D g  s			D g  gap-ofs			D g  aft-ofs			D  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??	 ?	??	 ?	-??	 ?	??	# ?	??	, ?	??	1 ?	??	B ?	??	D ?	?? 		D	  g  nameg  	point+-n!? CNRNM.01 	     h?   ?  ]*4 5?$  4 5"  $  "  ?$  4 5"  $  "  Y4 5?
?$  "  74
?$  "   4 54 54 5>  "  G  C    ?      g  gb
	 ? g  	new-point	 ? g  pmax			 ? g  t		  ? g  t		C ? g  delta		Y ? g  t		_ ?  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	 ??	  ?	??	1 ?	??	5 ?	??	6 ?	 ??	C ?	??	R ?	"??	Y ?	??	Y ?	??	_ ?	??	_ ?	
??	l ?	??	p ?	??	t ?	?? ? ?	?? ? ?	 ?? ? ?	.?? ? ?	?? 	 ?	  g  nameg  gb-goto-char? CR.OP01       h0     ]	4 54
4 5544 556    
      g  gb
		, g  s			,  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	"??	 ?	??	 ?	??	" ?	 ??	* ?	??	, ?	?? 		,  g  nameg  
gb->string? CR h0     ]
44 554 >  "  G   6            g  gb
		* g  string-proc		* g  new			*  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	* ?	?? 		*	  g  nameg  
gb-filter!? CRQPC      hX   |  ]!4 5"  <4
5$  ?45?"???45?6
"???   t      g  gb
		U g  str			U g  start			K g  acc			K g  t			K  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	% ?	??	& ?	%??	3 ?	??	; ?	??	> ?	!??	I ?	??	K ?	??	K ?	??	M ?	??	U ?	?? 		U  g  nameg  	gb->lines? CR      h8   '  ]
44 554 >  "  G   4
56             g  gb
		1 g  
lines-proc		1 g  	new-lines			1  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
 
??		??		??		??		??		??	'	??	1	?? 			1	  g  nameg  gb-filter-lines!? CRGRS       h   ?   ]L  6      ?       g  c
		
  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?

	??	

	?? 		
   C     h   ?   ]L  6      ?       g  s
		
  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
	??	
	?? 		
   C01/.TJ56 	      h?   m  ] 4L 54L 5 4L 5?$  C4L 5454 >  "  G  4L  ?>  "  G  4L ?>  "  G  C     e      g  gap-ofs
		{ g  aft-ofs		{ g  s		+	{ g  c		6	{  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?
	??		??	
	??		??		 ??		??	"	??	%	??	+	??	.	??	6	??	9	??	O	??	V	$??	[	??	d	??	k	$??	p	?? 		{
   CU 	     hP     ]	4 5$  "  4 >  "  G   O  O  O  6          g  gb
		L g  t			,  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?

??		??			??		??		??	!	??	H		??	J	??	L	?? 
		L  g  nameg  make-gap-buffer-port? CRC?      g  m
		,  g  filenamef  X/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/gap-buffer.scm?		Y
??	-	l	??	3	l	??	5	m	??	7	l	??	:	k
??	;	s	??	H	s
??	I	u	??	S	u	-??	U	u	??	X	u
??	Y	v	??	c	v	-??	e	v	??	h	v
??	i	w	??	s	w	-??	u	w	??	x	w
??	y	x	?? ?	x	-?? ?	x	?? ?	x
?? ?	z	?? ?	z	-?? ?	z	?? ?	z
?? ?	{	?? ?	{	-?? ?	{	?? ?	{
?? ?	|	?? ?	|	-?? ?	|	?? ?	|
?? ?	}	?? ?	}	-?? ?	}	?? ?	}
?? ? ?
?? ? ?
?? ? ?
??? ?
??? ?	??? ?	+??? ?	??? ?
??? ?
??
x ?
??^ ?
??( ?
??N ?
??; ?
??? ?
??I ?
?? ?
??0 ?
??? ?
??~ ?
??K ?
??? ?
?? ?
??? ?
??!r 
??&?
?? C	&?
   C6 