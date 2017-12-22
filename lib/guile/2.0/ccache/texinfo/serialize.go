GOOF----LE-4-2.0ªK      ]  4   h»      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  texinfo¤	g  	serialize¤	 ¤		g  filenameS¤	
f  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm¤	g  importsS¤	 ¤	 ¤	g  string-utils¤	 ¤	 ¤	g  sxml¤	g  	transform¤	 ¤	 ¤	g  srfi¤	g  srfi-1¤	 ¤	 ¤	g  srfi-13¤	 ¤	 ¤	 ¤	g  exportsS¤	g  stexi->texi¤	 ¤	 g  set-current-module¤	!  ¤	"  ¤	#g  reverse¤	$g  list-intersperse¤	%g  reverse!¤	&g  filter*¤	'g  append¤	(g  list*¤	)f  
¤	*f   ¤	+f  @¤	,g  include¤	-g  empty-command¤	.g  string=?¤	/f  *braces*¤	0f  @}¤	1g  
append-map¤	2f  @{¤	3f  }¤	4f  {¤	5g  inline-text¤	6g  map¤	7g  warn¤	8f  Strange inline-args!¤	9g  error¤	:f  Invalid inline-args¤	;f   ¤	<g  
drop-while¤	=g  not¤	>g  assq-ref¤	?f  ,¤	@g  inline-args¤	A? ¤	Bg  inline-text-args¤	C* ¤	Dg  serialize-text-args¤	Eg  eol-text-args¤	Fg  eol-text¤	Gf  , ¤	Hg  eol-args¤	If  @bye
¤	Jf  
@c %**end of header

¤	Kg  title¤	Lf  
@settitle ¤	Mg  and=>¤	Ng  filename¤	Of  @setfilename ¤	Pf  8\input texinfo   @c -*-texinfo-*-
@c %**start of header
¤	Qf  

¤	Rf  @end ¤	Sg  
string-ref¤	Tg  string-length¤	Ug  environ¤	Vg  symbol->string¤	Wg  table-environ¤	Xg  fill-string¤	Yg  string-concatenate¤	Zg  
line-widthS¤	[g  break-long-words?S¤	\g  wrap¤	]g  	paragraph¤	^f  @item
¤	_g  item¤	`f  @item ¤	ag  entry¤	bf  
@c %end of fragment
¤	cf  
@c %start of fragment

¤	dg  fragment¤	eg  EMPTY-COMMAND¤	fg  INLINE-TEXT¤	gg  INLINE-ARGS¤	hg  INLINE-TEXT-ARGS¤	ig  EOL-TEXT¤	jg  EOL-TEXT-ARGS¤	kg  INDEX¤	lg  EOL-ARGS¤	mg  ENVIRON¤	ng  TABLE-ENVIRON¤	og  ENTRY¤	pg  ITEM¤	qg  	PARAGRAPH¤	rg  FRAGMENT¤	sg  serializers¤	tf  Unknown command type¤	u}{@ ¤	vg  escaped-chars¤	wg  string¤	xg  memq¤	yg  string->list¤	zg  escape¤	{g  string-concatenate-reverse¤	|g  string?¤	}g  assq¤	~g  texi-command-specs¤	g  symbol?¤ g  %¤ f  $Unknown stexi command, not rendering¤ f  Invalid stexi¤C 5hÈ@    ]4	
5 4" >  "  G   #      h@   _  ] (   C"   (  6"ÿÿà   "ÿÿÏW      g  src-l
		@ g  elem		@ g  l			/ g  dest			/  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	"
		#			$			%			%			&		!	&	 	&	&	(	'	&		/	&		/	$		2	$		5	$	-	8	$	9	@	$	 		@	  g  nameg  list-intersperse C$R%        hh     ]"  U(  6$  #4 5$  
"  "ÿÿÊ4 5$  	"  "ÿÿ«"ÿÿ         g  pred
		f g  l		f g  in			[ g  out			[  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	)
		*			+			,			-			+			.			.		"	.	"	$	.		(	.		+	.	2	.	.	,	<	.		=	0		>	0		H	0		M	0	!	[	0		[	*		^	*		f	*	 		f	  g  nameg  filter* C&R#'   hx   ¾  -  1  3 4 5"  K(  C$  45"ÿÿÚ(  "ÿÿÆ"ÿÿµ"ÿÿ¨     ¶      g  args
			s g  args		s g  tail			s g  in			f g  out			f  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	3
	
	4			4			5			4			6		!	7		'	8		(	8		,	7		/	8	"	0	8	+	5	8	3	9	8	+	A	8		D	9		H	7		K	9	"	U	9		X	:		[	:	%	^	:		f	:		f	6		i	6		s	6	 			s


  g  nameg  list* C(R()$*+ h    [  ]456       S      g  exp
		 g  lp		 g  command			 g  type			 g  formals			 g  args			 g  accum			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	?			@			A			C			A			D			D			@	 				  g  nameg  include C,R(*+        h   @  ]6  8      g  exp
		 g  lp		 g  command			 g  type			 g  formals			 g  args			 g  accum			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	F
		G			
	G			G	 			  g  nameg  empty-command C-R./(01h   ¬   ]L  6       ¤       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	O			O	+			O	% 			   C#23    h   ¬   ]L  6       ¤       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	L			L	+			L	% 			   C4+     hP   Î  ]45$  4O 4 556	4
O 4 556 Æ      g  exp
		O g  lp		O g  command			O g  type			O g  formals			O g  args			O g  accum			O  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	I
		J				J			J			J			N			O			O	1	#	O	:	%	O	1	'	O		)	P		-	N		1	K		2	L		<	L	1	A	L	:	C	L	1	E	L		G	M		K	M		O	K	 		O	  g  nameg  inline-text C5R(3$6789:;        h@     ] $  3 $  $ $  4L >  "  G  "    CL 6C  ü       g  x
		>  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	W			X			Y			X			Z	#		Z			Z			[			[	"	#	[		2	\		7	]	$	;	]		=	X	  		>   C<=>  h   ¤   ]L  6             g  x
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	_	 	
	_	, 		
   C#?;4+       hP   ¦  ]$  144O 44O 4	5555
5"  6             g  exp
		I g  lp		I g  command			I g  type			I g  formals			I g  args			I g  accum			I  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	R
		S				T				U			V			^			_		)	`	 	1	_		3	^		5	V		7	a		9	U		?	T		A	b			E	b		I	S	 		I	  g  nameg  inline-args C@R(3'$61       h   ¬   ]L  6       ¤       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	k	(		k	:			k	4 			   C#       h   ¬   ]L O 4 56    ¤       g  x
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	k			k	@		k	 		   C<=>  h   ¤   ]L  6             g  x
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	m	!	
	m	- 		
   C#A;4+       hP   »  ]$  6444O 44	O 4
55555?"  6  ³      g  exp
		N g  lp		N g  command			N g  type			N g  formals			N g  args			N g  accum			N  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	d
		e				f				g			i			j			l		"	m		,	n	!	4	m		6	l		8	j		:	o		<	i		>	g		D	f		F	p			J	p		N	e	 		N	  g  nameg  inline-text-args CBR'$61      h   ¬   ]L  6       ¤       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	v	#		v	5			v	/ 			   C h   ¦   ]L O  6        g  arg
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	v				v	 		   C#<=>      h   ¤   ]L  6             g  x
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	z		
	z	  		
   CC 
      h@   K  ]44 O 444O 455555	5@      C      g  lp
		: g  formals		: g  args			:  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	r
		u			v			w				y	
		z		&	{		.	z		0	y	
	2	w			4	v		6	|		8	u		:	s	 		:	  g  nameg  serialize-text-args CDR()D*+    h    S  ]456     K      g  exp
		 g  lp		 g  command			 g  type			 g  formals			 g  args			 g  accum			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
	~
					 			 			 				 			  g  nameg  eol-text-args CER()1 h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 		 	'		 	! 			   C#*+        h8     ]4O 4$   "   556            g  exp
		2 g  lp		2 g  command			2 g  type			2 g  formals			2 g  args			2 g  accum			2  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 
	 			 			 		 		 	'	$ 	2	& 		( 			* 			. 		2 	 		2	  g  nameg  eol-text CFR()$'<=6> h   ¦   ]L  6             g  x
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 	"	
 	. 		
   C#G*+       h8   ¢  ]4444O 4	555?
56        g  exp
		6 g  lp		6 g  command			6 g  type			6 g  formals			6 g  args			6 g  accum			6  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 
	 			 			
 	
	 		 		 	"	$ 		& 		( 	
	* 	
	, 			. 			2 		6 	 		6	  g  nameg  eol-args CHR(I1h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 		 	*		 	$ 			   C#J>KLMN)#O h   Ú   ]4 5CÒ       g  filename
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 		 		 	$		 	3	 	-	 	$	 	 		   C;PQR    h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 ¢	$	 ¢	6		 ¢	0 			   CST)D*+  hØ   K  ] $  T4O 4 5544	55
4455$  "  64O 4$   "   55(  "  4455
$  "  456
C      g  exp
	 Ø g  lp	 Ø g  command		 Ø g  type		 Ø g  formals		 Ø g  args		 Ø g  accum		 Ø g  key		 Ø g  t		K	^ g  body	  Á  
g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 
	 		 		 		 		 	0	$ 	9	' 	0	) 		+ 		, 		/ 		5 	$	7 		9 		; 	-	< 		? 		E 	&	G 		K 		K 		[ 		` 		d 		h ¡		l ¡		m ¢		w £	$	 £	-  £	6  £	A  £	$  ¢	  ¢	  ¤	  ¤	  ¥	 ¡ ¥	( ¢ ¦	, § ¦	; © ¦	, ª ¦	( ¬ ¥	 ¯ ¥	 ³ ¤	 » ©	 ¾ ©	 Ã ª	 Ä «	 Ð ¬	 Ô ¬	 Ø ¡	 6	 Ø	  g  nameg  environ CUR(QR1       h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 °		 °	'		 °	! 			   C#);V+*  hh     ]4O 4$   "   55$  	"  $  4	5
 "  
6

      g  exp
		h g  lp		h g  command			h g  type			h g  formals			h g  args			h g  accum			h g  arg		?	^  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 ®
	 ¯			
 ¯		 °			 ±		 ±		  ±	'	( ±	2	* ±		, °			. ²			4 ³		7 ³		? ³	+	? ³			D ´		H ´		I µ		N µ	%	P µ		R µ	0	U µ		` ·			d ·		h ¯	 		h	  g  nameg  table-environ CWRXYZ[       h   Å   ]4 5	H6    ½       g  strings
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 ¹
	 º		 º	 		  g  nameg  wrap C\R(Q\#1   h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Â		 Â	)		 Â	# 			   C      h(     ]444O 4 55556      g  exp
		( g  lp		( g  command			( g  type			( g  formals			( g  args			( g  accum			(  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 ¾
	 ¿			 À			
 Á	
	 Â		 Â	/	 Â	8	 Â	/	  Â		" Á	
	$ À			( ¿	 		(	  g  nameg  	paragraph C]R(1      h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Æ		 Æ	'		 Æ	! 			   C#^  h    _  ]4O 4 556  W      g  exp
		 g  lp		 g  command			 g  type			 g  formals			 g  args			 g  accum			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Å
	 Æ			 Æ	-	 Æ	6	 Æ	-	 Æ			 Ç			 Æ	 				  g  nameg  item C_R(1       h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Ë		 Ë	'		 Ë	! 			   C#)   h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Í		 Í	'		 Í	! 			   C`    h8     ]4O 4 554O 4556        g  exp
		6 g  lp		6 g  command			6 g  type			6 g  formals			6 g  args			6 g  accum			6  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Ê
	 Ë			 Ë	-	 Ë	6	 Ë	-	 Ë			 Ì			 Í			& Í	-	+ Í	6	. Í	-	0 Í			2 Î			6 Ë	 		6	  g  nameg  entry CaR(b1      h   ¯   ]L  6       §       g  x
			  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Ó		 Ó	'		 Ó	! 			   C#c  h    l  ]4O 4 556d      g  exp
		  g  lp		  g  command			  g  type			  g  formals			  g  args			  g  accum			   g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 Ñ
	 Ò			 Ó			 Ó	-	 Ó	6	 Ó	-	 Ó			 Ô			  Ò	 
		 	  g  nameg  fragment CdRe-if5ig@ihBiiFijEikEilHimUinWioaip_iq]irdi,i sR>s9t       h8   j  ]45$  "  4 5 6    b      g  exp
		4 g  lp		4 g  command			4 g  type			4 g  formals			4 g  args			4 g  accum			4 g  t			4  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 è
	 é		 é		 ê		 ê		$ ê		4 é	 		4	  g  nameg  	serialize CRuvRw#xvy        h`   Ý  ]"  D(  45@45$  @"ÿÿÍ"ÿÿ¼4 5"ÿÿ¬      Õ      g  str
		Z g  in		J g  out			J  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 î
	 ð		 ñ		 ò		 ò		 ó		 ó		! ó		% ó		( ô		+ ô	"	0 ô		9 ô		< õ		? õ	!	B õ		J õ		J ð		K ð		R ð	*	Z ð	 		Z  g  nameg  escapeg  documentationf  BEscapes any illegal texinfo characters (currently @{, @}, and @@). CzR{|z}~V&79       hØ   d  ]
 $   "  $  C4 5$  4 5C $  4 5$  r L 4 545"  $  "  @ "  9 $  - $   	&  	 "  "ÿÿÃ"  "ÿÿ»"  "ÿÿ³64
 >  "  G  C 6 \      g  in
	 × g  out	 × g  command-spec		> Ï  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 ú		 ü		 ü		 ü		 û		 ý		" û		# ý		, ý		0 þ		4 û		5 ÿ		: ÿ	!	> ÿ		> ÿ		F 			M		R	(	T		W		Y		`	)	c		j		o			w		{			~
	&	
	 		 
	7 
	0 
	 	$ 	/ 
	 	 ¸	 ¹	 ½	 Ä	 Ó	 ×	 ,	 ×	  g  nameg  lp C h    2  ]	O Q 4 56  *      g  tree
		 g  lp		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm
 ÷
	 ú		 ú		 ú		 ù	 		  g  nameg  stexi->texig  documentationf  2Serialize the stexi @var{tree} into plain texinfo. CRC x      g  m
		,  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/texinfo/serialize.scm		
â	"
í	)
9	3
Î	?		;	F
	I
Y	R
þ	d
	r
	~
B 
! 
'þ 
+m ®
,h ¹
/ ¾
1j Å
4ñ Ê
7g Ñ
7i Ø	7Õ ×
9 è
9 í	9 í
;ø î
@Å ÷
 	@Ç
   C6 