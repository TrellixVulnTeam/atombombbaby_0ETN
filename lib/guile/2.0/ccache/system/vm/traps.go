GOOF----LE-4-2.0Ï      ] Ú 4    h      ] g  guile¤	 ¤	g  define-module*¤	 ¤	 ¤	g  system¤	g  vm¤	g  traps¤		 ¤	
g  filenameS¤	f  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm¤	g  importsS¤	g  base¤	g  pmatch¤	 ¤	 ¤	 ¤	 ¤	g  frame¤	 ¤	 ¤	g  program¤	 ¤	 ¤	g  objcode¤	 ¤	 ¤	g  instruction¤	 ¤	 ¤	g  xref¤	  ¤	!  ¤	"g  rnrs¤	#g  bytevectors¤	$"# ¤	%$ ¤	&!% ¤	'g  exportsS¤	(g  trap-at-procedure-call¤	)g  trap-in-procedure¤	*g  trap-instructions-in-procedure¤	+g  trap-at-procedure-ip-in-range¤	,g  trap-at-source-location¤	-g  trap-frame-finish¤	.g  trap-in-dynamic-extent¤	/g  trap-calls-in-dynamic-extent¤	0g  #trap-instructions-in-dynamic-extent¤	1g  trap-calls-to-procedure¤	2g  trap-matching-instructions¤	3()*+,-./012 ¤	4g  set-current-module¤	54 ¤	64 ¤	7g  make-syntax-transformer¤	87 ¤	97 ¤	:g  	arg-check¤	;g  macro¤	<g  $sc-dispatch¤	=< ¤	>< ¤	?g  any¤	@???? ¤	Ag  syntax-object¤	Bg  if¤	Cg  top¤	DC ¤	Eg  ribcage¤	Fg  dummy¤	Gg  arg¤	Hg  
predicate?¤	Ig  message¤	JFGHI ¤	Kg  m-168148b5-0¤	LKC ¤	MLDDD ¤	Nf  l-168148b5-5¤	Of  l-168148b5-6¤	Pf  l-168148b5-7¤	Qf  l-168148b5-8¤	RNOPQ ¤	SEJMR ¤	TE ¤	Ug  x¤	VU ¤	WL ¤	Xf  l-168148b5-2¤	YX ¤	ZEVWY ¤	[DSTZ ¤	\g  hygiene¤	]\ ¤	^AB[] ¤	_g  not¤	`A_[] ¤	ag  error¤	bAa[] ¤	cf  bad argument ~a: ~a¤	dAc[] ¤	eg  quote¤	fAe[] ¤	g??? ¤	hFGH ¤	iLDD ¤	jf  l-168148b5-e¤	kf  l-168148b5-f¤	lf  l-168148b5-10¤	mjkl ¤	nEhim ¤	oDnTZ ¤	pABo] ¤	qA_o] ¤	rAao] ¤	sf  bad argument ~a: expected ~a¤	tAso] ¤	uAeo] ¤	vg  syntax-violation¤	wv ¤	xv ¤	yf  -source expression failed to match any pattern¤	zf  trap already enabled¤	{f  trap already disabled¤	|g  new-disabled-trap¤	}g  new-enabled-trap¤	~g  	procedure¤	g  frame-procedure¤ g  program?¤ g  program-objcode¤ g  frame-matcher¤ g  vmS¤ 	¤ g  closure?S¤ 	¤ g  
our-frame?S¤ 	¤  ¤ g  the-vm¤ g  
procedure?¤ g  proc¤ g  handler¤ g  	add-hook!¤ g  vm-apply-hook¤ g  remove-hook!¤ g  current-frameS¤ 	¤ 	¤ 	¤ 	¤  ¤ g  enter-handler¤ g  exit-handler¤ g  warn¤ f  already in proc¤ g  frame-previous¤ g  vm-push-continuation-hook¤ g  vm-pop-continuation-hook¤ g  vm-abort-continuation-hook¤ g  vm-restore-continuation-hook¤  g  next-handler¤ ¡g  vm-next-hook¤ ¢g  number?¤ £g  integer?¤ ¤g  exact?¤ ¥g  non-negative-integer?¤ ¦g  positive-integer?¤ §g  and-map¤ ¨g  range?¤ ©g  or-map¤ ªg  	in-range?¤ «g  range¤ ¬g  frame-address¤ ­g  frame-instruction-pointer¤ ®g  objcode->bytecode¤ ¯g  program-last-ip¤ °f  
unexpected¤ ±g  for-each¤ ²g  	assv-set!¤ ³g  assv-ref¤ ´g  sort!¤ µg  program-sources-pre-retire¤ ¶g  program-sources-by-line¤ ·f  no instructions found at¤ ¸f  :¤ ¹f  ; using line¤ ºf  instead¤ »f  no instructions found for¤ ¼g  source->ip-range¤ ½g  source-closures¤ ¾g  source-procedures¤ ¿g  source-closures-or-procedures¤ À ¤ Ág  string?¤ Âg  file¤ Ãg  	user-line¤ Äg  map¤ Åf  No procedures found at ~a:~a.¤ Æ	¤ ÇÆ ¤ Èg  frame?¤ Ég  return-handler¤ Êg  abort-handler¤ Ëf  .return-or-abort traps may only be enabled once¤ Ì	¤ Í	¤ Î	¤ Ï	¤ ÐÌÍÎÏ ¤ Ñg  apply-handler¤ Ò	¤ Ó	¤ Ô	¤ ÕÒÆÓÔ ¤ Ög  length¤ ×g  delq¤ Ø ¤ Ùg  
frame-pred¤C 5   hàx  ]  ]4	
&'35 46 >  "  G   49:;>@^`bdf  h(   k   ]     C    c       g  dummy
		$ g  arg		$ g  
predicate?			$ g  message			$  		$	   Cgpqrtu     h0   V   ]      C       N       g  dummy
		) g  arg		) g  
predicate?			)  		)	   Cxy   h@   $  ]4 5$  @4 5$  @ 6             g  x
		9 g  tmp		9 g  tmp		"	9  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	N	 		9  g  
macro-typeg  syntax-rulesg  patternsg  argg  
predicate?g  message g  argg  
predicate?   C5:Raza{ hP   è   -  . , 3  #   M$  "  4>  "  G  4L >  "  G  NL C       à       g  frame
		I  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	c			d		 	d		$	d		)	d		2	e		D	f	 		I
  g  nameg  disable-trap C     hX   ç   -  . , 3  #   M$  4>  "  G  "   4L >  "  G  NLL LO C      ß       g  frame
		R  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	]			^			^		 	^		%	^		2	_		F	`	 		R
  g  nameg  enable-trap C      h      ]HO Q C        g  vm
		 g  enable		 g  disable			 g  enabled?			 g  enable-trap			  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	V
		W	 			  g  nameg  new-disabled-trap C|R|       h     ]4 56       û       g  vm
		 g  frame		 g  enable			 g  disable			  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	k
		l			l	 			  g  nameg  new-enabled-trap C}R~     h8   ý   ]	4 5L &  C45$  454L 5CC      õ       g  frame
		2 g  
frame-proc			2  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	s			t				t	
		u			u			v		 	v		!	w		(	x		/	w	 		2   C     h   ®   ]4 5L C   ¦       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	y			z			z	
 		   C       h0     ]
 ¦$  4 5"   $  O CO C        g  proc
		. g  match-objcode?		. g  proc			.  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	n
		o				o		
	p			o			r	 		.	  g  nameg  frame-matcher CRas    h   Ñ   ]4L 5$  L  6C   É       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 		 	
	 		 	
 		  g  nameg  
apply-hook C}     h   ±   ]4L5L 6 ©       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 		 		 	 		   C       h   ±   ]4L5L 6 ©       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 		 		 	 		   C    h°     - /   0   3 #  45 #  #  4 54 5$  "  4>  "  G  45$  "  4	>  "  G  
O Q O O 6         g  proc
	 ­ g  handler	 ­ g  vm		 ­ g  closure?		 ­ g  
our-frame?		 ­ g  
apply-hook	  ­  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	~
		~	8	0 	-	; 		` 	 ­ 	 	 ­	
g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  trap-at-procedure-call C(Ras       hX     ]M$  4L >  "  G  N"   4L  5$  %M$   64L >  "  G  NCC           g  frame
		S  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 ³		 ´			 ¯		 °		" ¶	
	, ¶		2 ¦		6 §		: §	
	; ©		O ª	 		S  g  nameg  
apply-hook C      h(   Õ   ]M$  4L  >  "  G  NCC       Í       g  frame
		!  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 ¹		 º			 ¯		 °	 		!  g  nameg  push-cont-hook C hh   I  ]	M$  4L >  "  G  N"   4L 4 55$  .4 5M$  64L>  "  G  NCC       A      g  frame
		a g  frame	8	_  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 ½		 ¾			 ¯		 °		" À	
	% À		- À	
	1 À		2 Á		8 Á	
	@ ¦		D §		H §	
	I ©		] ª	 		a  g  nameg  pop-cont-hook C       hX     ]M$  4L >  "  G  N"   4L  5$  %M$   64L >  "  G  NCC           g  frame
		S  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 Ã		 Ä			 ¯		 °		" Æ	
	, Æ		2 ¦		6 §		: §	
	; ©		O ª	 		S  g  nameg  
abort-hook C        hX     ]M$  4L >  "  G  N"   4L  5$  %M$   64L >  "  G  NCC     
      g  frame
		S  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 É		 Ê			 ¯		 °		" Ì	
	, Ì		2 ¦		6 §		: §	
	; ©		O ª	 		S  g  nameg  restore-hook C} 	       hÀ   x  ]44L5L>  "  G  44L5L>  "  G  44L5L>  "  G  44L5L>  "  G  44L5L>  "  G   $  24L 5$  %M $   64L >  "  G  N CCC       p      g  frame
	 ¹  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 Ñ		 Ò		 Ò		 Ò		 Ó		 Ó		, Ó		5 Ô		8 Ô		E Ô		N Õ		Q Õ		^ Õ		g Ö		j Ö		w Ö	  ×	  ×	  ×	  ¦	  §	  §	
  ©	 ³ ª	 	 ¹   C        h   <  ]M$  4L >  "  G  N"   44L5L>  "  G  44L5L>  "  G  44L5L>  "  G  44L5L>  "  G  4L5L 6      4      g  frame
	   g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 Ù		 Ú			 ¯		 °		" Ü		% Ü		2 Ü		; Ý		> Ý		K Ý		T Þ		W Þ		d Þ		m ß		p ß		} ß	  à	  à	 	    C hX  }  - /   0   3 #  #  45 #  #  4 54 5$  "  4>  "  G  45$  "  4	>  "  G  45$  "  4
>  "  G  HO O 	O 
O O Q Q 	Q 
Q Q 
	O 	
	O 6       u      g  proc
	Q g  enter-handler	Q g  exit-handler		Q g  current-frame		Q g  vm		Q g  closure?		Q g  
our-frame?		Q g  in-proc?	 ³Q g  
apply-hook	 êQ g  push-cont-hook		 êQ g  pop-cont-hook	
 êQ g  
abort-hook	 êQ g  restore-hook	 êQ  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 
	# 	4	9  	(	D ¡		i ¢	  £	 ³ ¤	Q Ï	 		Q	
g  current-frameS	g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  trap-in-procedure C)Ras  h   Ð   ]4L 5$  L  6C   È       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 í		 î	
	 î		 ï	
 		  g  nameg  	next-hook C)¡      h8   ð   ]44L5L>  "  G   $  4L 5$  L  6CC  è       g  frame
		6  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 ñ		 ò		 ò		 ò		! ó		" î	
	, î		2 ï	
 			6  g  nameg  enter C¡        h(   Ë   ]4L >  "  G  4L 5L6       Ã       g  frame
		!  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 õ		 ö		 ÷		! ÷	 		!  g  nameg  exit C    hð   õ  - /   0   3 #  #  45 #  #  4 54 5$  "  4>  "  G  45$  "  4	>  "  G  45$  "  4
>  "  G  O Q  O O 6	    í      g  proc
	 ì g  next-handler	 ì g  exit-handler		 ì g  current-frame		 ì g  vm		 ì g  closure?		 ì g  
our-frame?		 ì g  	next-hook	 » ì  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 ä
	# å	A	9 è	*	D é		i ê	  ë	 ì ù	 	 ì	
g  current-frameS	g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  trap-instructions-in-procedure C*R¢£¤      h0     ]4 5$   4 5$  4 5$   
CCCC ü       g  x
		/  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 ý
	 þ		 þ		 þ		 þ		 þ	 	# þ		& þ	0	( þ	+ 
		/  g  nameg  non-negative-integer? C¥R¢£¤       h0   ÷   ]4 5$  4 5$  4 5$   
CCCC  ï       g  x
		.  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
 
										 	#		&	+ 			.  g  nameg  positive-integer? C¦R§¥     h    ã   ] $  4 5$   6CCÛ       g  x
		   g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
							
			.						.		 
		    C  h   È   ] $   6C     À       g  x
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm

							 		  g  nameg  range? C¨R©h   Í   ] L $  L  CC   Å       g  bounds
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
	
										 		   C        h   Ì   ]O  6 Ä       g  range
		 g  i		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm

		 			  g  nameg  	in-range? CªRas¨«¨      h@   +  ]	"  ,"  N C$   $  	"ÿÿÜ"ÿÿÜ"ÿÿØM "ÿÿÌ      #      g  fp
		: g  frames		2  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
				!									#		 	!		$ 		* 		2	 		:  g  nameg  cull-frames! C¬­ª   hp   ¯  ]!4 54 54L>  "  G  4L5M $  
M "  $  $  CM N C$  M N L 6C §      g  frame
		o g  fp		o g  ip			o g  now-in-range?		A	o g  was-in-range?		A	o  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
#		$		
%		$		&		''		2(	"	6(		9(	6	<(	3	A'		K)	
	Q+		V,		X,		`)	
	e.		g.		m/	 		o  g  nameg  next-handler C*¬ h(      ]M $  M 4 5$  M N CCC     ø       g  frame
		#  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
1		2			2		3		3	!	3		2	
	4		4	
 
		#  g  nameg  exit-handler C       hø   3  - /   0   3 
#  #  45 #  #  4 54 5$  "  4>  "  G  4	5$  "  4
>  "  G  45$  "  4>  "  G  HO O 	Q Q 	 	O 6	     +      g  proc
	 ó g  range	 ó g  handler		 ó g  current-frame		 ó g  vm		 ó g  closure?		 ó g  
our-frame?		 ó g  fp-stack	 ³ ó g  cull-frames!	 È ó g  next-handler		 È ó  
g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm

	#	@	9	)	D		i	 	 ³	 ³	 ó6	 
	 ó	
g  current-frameS	g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  trap-at-procedure-ip-in-range C+R®  h   Ý   ]44 55
ºC       Õ       g  prog
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
=
	>		>	0	>		>	 		  g  nameg  program-last-ip C¯Ra°¯±²³       h0     ]	4M   4M  5$  "  5N C      g  pair
		0 g  t		)  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
T		V		
V	$	W	$	X	(	X	8	X	(	X	$	&Y	(	*W		,V		.U	 		0   C´     h   Å   ] C      ½       g  x
		
 g  y		
  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
[		[	(	[	0		[	% 		
	   Cµ 	     h   û  ]R" $  Ì$  «$  $  g	$  L	$  0

$  

"  45
"  
4 5"  	"  45"  45"  45"ÿÿ-H4O >  "  G  4J>  "  G  JC4 5"ÿþê    ó      g  proc
	 g  file	 g  sources		 g  out		 g  v		 Ñ g  vx		" Ã g  vy		" Ã g  vx		3 ² g  vy		3 ² g  vx			A ¡ g  v	
	^  g  vx		j	x g  alist	 Ú  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
@
	A			C		C		D		E		E		LG		PG		UI	*	VI	#	ZI		]J	+	^J	#	uL	&	}M	+ M	2 M	+ N	2 N	# H	 H	 ¦Q	 ªQ	 ¬Q	 »Q	 ½Q	 ÌQ	 ÎQ	 ÙD	 ÚR	 ÚR	 ÝS	
 ö[	
A	A	B	A	 '		  g  nameg  program-sources-by-line C¶R©·¸¹º h@   0  ] L$   C L$  #4L L >  "  G   CC (      g  line-and-ranges
		?  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
_		a		a		`		b		c		c		`		d		d		#d	8	'e		*e	'	,e	=	1d		<f	 		?   C¶»¸     h@   >  ]4O 4 55$  C4>  "  G  C  6      g  proc
		> g  file		> g  line			> g  t			>  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
^
	_		h		_		_		%j		)j		-j	/	4j		=k	 		>	  g  nameg  source->ip-range C¼R½¾       h(   .  ]
4 5$  D4 5D   &      g  file
		% g  line		% g  closures			%  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
m
	n		n		o		o		p		q		%q	 			%	  g  nameg  source-closures-or-procedures C¿RÀÁasÂÁ¦Ã¦¿}Ä¼+       h(   Ý   ]	4 L L5 LLLL6	  Õ       g  proc
		& g  range		&  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
			!		=		!			&	 		&   CaÅ    h0   Ì   ]4LLLLLLO L5N M (  
LL6C  Ä       g  frame
		.  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
							"		&		,	 		.   C±       h   §   ] L 6       g  trap
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
			$ 		   C      h    ±   ]4 O M >  "  G  N C©       g  frame
		   g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
						 		    C    hØ   æ  - /   0   3 #  #  45 4 5$  "  4>  "  G  45$  "  4	
>  "  G  45$  "  4>  "  G  H4 >  G  O O 6    Þ      g  file
	 Ô g  	user-line	 Ô g  handler		 Ô g  current-frame		 Ô g  vm		 Ô g  traps	  Ô g  procs	 © Ô g  	closures?	 © Ô  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
w
	#x	:	*y		Oz		t{	 |	 ~	 £~	7 ¨~	 ¬}	 Ô	 	 Ô	
g  current-frameS	g  vmS	   g  nameg  trap-at-source-location C,RÇÈasÈÉÊ¬¬        h(   æ   ]M$  4 5M&  	NL  6CC      Þ       g  frame
		"  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
								
	¡		¢	 		"  g  nameg  pop-cont-hook C¬    h(   ì   ]M$  4 5M$  	NL  6CC     ä       g  frame
		#  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
¤		¥			¥		¥		¥	
	§		¨	 		#  g  nameg  
abort-hook C}aË  h`     ]M$  "  4>  "  G  44L5L>  "  G  44L5L >  "  G  4L5L 6         g  frame
		]  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
¬		­		®		®		®		¯		"¯		/¯		8°		;°		H°		S±		]±	 		]   C hH   ð   ]N44L5L>  "  G  44L5L >  "  G  4L5L 6    è       g  frame
		D  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
²		³		´			´		´		µ		"µ		/µ		:¶		D¶	 		D   C     hà   £  - /   0   3 #  45 4 5$  "  4>  "  G  45$  "  4	
>  "  G  45$  "  4
>  "  G  4 5HO O Q Q  O O 6            g  frame
	 Ú g  return-handler	 Ú g  abort-handler		 Ú g  vm		 Ú g  fp	  Ú g  pop-cont-hook	 ª Ú g  
abort-hook	 ª Ú  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm

		&	!		F		k	 	 	 Úª	 		 Ú	
g  vmS	   g  nameg  trap-frame-finish C-RÐasÉÊ h    Ñ   ]4M >  "  G  NL  6   É       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
É		Ê		Ë		Ì	 		  g  nameg  
abort-hook C-  h    Ò   ]4M >  "  G  NL  6   Ê       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
Ä		Å		Æ		Ç	 		  g  nameg  return-hook C hH   ì   ]M$  "  4L 5$  ,4L  >  "  G  4 LLO LL5NCC  ä       g  frame
		F  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
Î		Ï	
	Ï		Ï		Ñ		+Ó		BÒ	 		F  g  nameg  
apply-hook C}  h   ±   ]4L5L 6 ©       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
Ø		Ù		Ù	 		   C       hH   Þ   ]M$  +4M >  "  G  N4L >  "  G  "   N4L 5L6     Ö       g  frame
		C  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
Ú		Û			Ê		Ë		Ì		6Ý		9Þ		CÞ	 			C   C       h(  T  - /   0   3 #  #  45 #  #  4 54 5$  "  4>  "  G  45$  "  4	>  "  G  45$  "  4
>  "  G  45$  "  4>  "  G  HO 	O 
Q 		Q 

O 
O 6    L      g  proc
	$ g  enter-handler	$ g  return-handler		$ g  abort-handler		$ g  current-frame		$ g  vm		$ g  closure?		$ g  
our-frame?		$ g  	exit-trap	 Ø$ g  
abort-hook		 ð$ g  
apply-hook	
 ð$  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
»
	#¼	9	9¾	-	D¿		iÀ	 Á	 ³Â	 ØÃ	$Ö	 
	$	
g  current-frameS	g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  trap-in-dynamic-extent C.RasÑÉ  h   È   ]M N C       À       g  frame
			  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
ì		í		í	 			  g  nameg  
trace-push C      h    Ð   ]4L M >  "  G  M N C   È       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
ï		ð		ñ		ñ	 		  g  nameg  	trace-pop C      h   À   ]L M 6      ¸       g  frame
		
  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
ó		
ô	 		
  g  nameg  trace-apply C.  hH   ù   ]44L5L>  "  G  44L5L>  "  G  4L5L 6       ñ       g  frame
		A  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
ø		ù		ù		ù		ú		ú		,ú		7û		Aû	 
		A  g  nameg  enter C   hH   ú   ]44L5L>  "  G  44L5L>  "  G  4L5L 6       ò       g  frame
		A  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
		þ		þ		þ		ÿ		ÿ		,ÿ		7 		A 	 
		A  g  nameg  return C  hH   ù   ]44L5L>  "  G  44L5L>  "  G  4L5L 6       ñ       g  frame
		A  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
		þ		þ		þ		ÿ		ÿ		,ÿ		7 		A 	 
		A  g  nameg  abort C      h   Q  - /   0   3 #  #  45 #  #  4 54 5$  "  4>  "  G  45$  "  4	>  "  G  45$  "  4
>  "  G  
HO O 	O 
Q Q 	Q 
 
	O 
	O 
	O 6
     I      g  proc
	 g  apply-handler	 g  return-handler		 g  current-frame		 g  vm		 g  closure?		 g  
our-frame?		 g  *call-depth*	 ³ g  
trace-push	 Ï g  	trace-pop		 Ï g  trace-apply	
 Ï  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
ã
	#ä	?	9ç	(	Dè		ié	 ê	 ³ë		 			
g  current-frameS	g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  trap-calls-in-dynamic-extent C/RÕas        h   ¿   ]L  6·       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
			 		  g  nameg  
trace-next C.¡       h   Ã   ]4L5L 6 »       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
					 		  g  nameg  enter C¡     h   Ä   ]4L5L 6 ¼       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
					 		  g  nameg  return C¡    h   Ã   ]4L5L 6 »       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
"					 		  g  nameg  abort C    hÈ   Ö  - /   0   3 #  #  45 #  #  4 54 5$  "  4>  "  G  45$  "  4	>  "  G  
O Q  O O O 6
 Î      g  proc
	 Ç g  next-handler	 Ç g  current-frame		 Ç g  vm		 Ç g  closure?		 Ç g  
our-frame?		 Ç g  
trace-next	  Ç  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm

	#	F	9	/	D		i	 Ç%	 	 Ç	
g  current-frameS	g  vmS	g  closure?S	g  
our-frame?S	   g  nameg  #trap-instructions-in-dynamic-extent C0RÇasÑÉÖ¬-×       h0   ä   ]4M >  "  G  4MM5NNL  L6      Ü       g  frame
		*  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
?		:		<		;		"=		*A	 		*  g  nameg  return-hook C×      h(   Ú   ]4M >  "  G  4MM 5N NC    Ò       g  frame
		$  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
D		:		<		;		"=	 		$  g  nameg  
abort-hook C h`   ;  ]4M54L >  "  G  4 5&  CH4 LLO LO L 5KJMNC    3      g  frame
		\ g  depth			\ g  finish-trap		.	\  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
2		3			3		5		 7		+7		.8		1H		SG		XJ		ZI	 		\  g  nameg  
apply-hook C±       h   ª   ] L 6¢       g  disable
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
_		_	$ 		   C    h    ²   -  . , 3  #   L4L  56ª       g  frame
		   g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
[		\	)	 \	
 		 
   C   hH   Þ   -  . , 3  #   4 O M>  "  G  N4L 5L O C    Ö       g  frame
		D g  trap	8	D  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
^		_		/a	#	1a		2b	&	8b	 		D
   C       h   È   ]L  LO C  À       g  trap
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
Y	 		  g  nameg  with-pending-finish-disablers C( hÐ   Ô  - /   0   3 #  45 4 5$  "  4>  "  G  45$  "  4>  "  G  45$  "  4	>  "  G  H
O O Q Q 4 56       Ì      g  proc
	 É g  apply-handler	 É g  return-handler		 É g  vm		 É g  pending-finish-traps	  É g  
apply-hook	 ¦ É g  with-pending-finish-disablers	 ¦ É  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
+
	,	,	!-		F.		k/	 0	 0	 »e	 Éd	 
	 É	
g  vmS	   g  nameg  trap-calls-to-procedure C1RØasÙ      h   Ð   ]4L 5$  L  6C   È       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
n		o	
	o		p	
 		  g  nameg  	next-hook C}¡      h   ±   ]4L5L 6 ©       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
t		u		u	 		   C¡       h   ±   ]4L5L 6 ©       g  frame
		  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
v		w		w	 		   C    h   @  - /   0   3 #  45 4 5$  "  4>  "  G  45$  "  4>  "  G  	O  Q 
O O 6     8      g  
frame-pred
	  g  handler	  g  vm		  g  	next-hook		s   g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm
i
	j	/	!k		Fl	 r	 	 	
g  vmS	   g  nameg  trap-matching-instructions C2RC   U      g  m
		,  g  filenamef  W/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/traps.scm		8
ë	V
	k
	n
	~
 8 
& ä
'× ý
) 
+#
-
5N
6X=
<Î@
?ñ^
Aam
H)w
Oö
Xï»
c\ã
i¹
t7+
xÛi
 	xÝ
   C6 