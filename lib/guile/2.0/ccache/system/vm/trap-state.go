GOOF----LE-4-2.0?J      ] ? 4    h?      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  system?	g  vm?	g  
trap-state?		 ?	
g  filenameS?	f  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?	g  importsS?	g  base?	g  syntax?	 ?	 ?	g  srfi?	g  srfi-1?	 ?	g  selectS?	g  fold?	 ?	 ?	 ?	 ?	g  traps?	 ?	 ?	g  trace?	 ?	 ?	 g  frame?	!  ?	"! ?	#g  program?	$# ?	%$ ?	&"% ?	'g  exportsS?	(g  	add-trap!?	)g  
list-traps?	*g  trap-enabled??	+g  	trap-name?	,g  enable-trap!?	-g  disable-trap!?	.g  delete-trap!?	/g  with-default-trap-handler?	0g  install-trap-handler!?	1g  add-trap-at-procedure-call!?	2g  add-trace-at-procedure-call!?	3g  add-trap-at-source-location!?	4g  #add-ephemeral-trap-at-frame-finish!?	5g  add-ephemeral-stepping-trap!?	6()*+,-./012345 ?	7g  set-current-module?	87 ?	97 ?	:g  
make-fluid?	;g  %default-trap-handler?	<g  warn?	=f  Trap with no handler installed?	>g  default-trap-handler?	?g  make-record-type?	@f  <trap-wrapper>?	Ag  index?	Bg  enabled??	Cg  trap?	Dg  name?	EABCD ?	Fg  <trap-wrapper>?	Gg  make-trap-wrapper?	Hg  record-predicate?	Ig  trap-wrapper??	Jg  make-procedure-with-setter?	Kg  record-accessor?	Lg  record-modifier?	Mg  trap-wrapper-index?	Ng  trap-wrapper-enabled??	Og  trap-wrapper-trap?	Pg  trap-wrapper-name?	Qf  <trap-state>?	Rg  handler?	Sg  next-idx?	Tg  next-ephemeral-idx?	Ug  wrappers?	VRSTU ?	Wg  <trap-state>?	Xg  make-trap-state?	Yg  trap-state??	Zg  trap-state-handler?	[g  trap-state-next-idx?	\g  trap-state-next-ephemeral-idx?	]g  trap-state-wrappers?	^g  trap-wrapper<??	_g  error?	`f  Trap already enabled?	ag  setter?	ba ?	ca ?	dg  enable-trap-wrapper!?	ef  Trap already disabled?	fg  disable-trap-wrapper!?	gg  append?	hg  add-trap-wrapper!?	ig  delq?	jg  remove-trap-wrapper!?	kg  trap-state->trace-level?	lf  )no wrapper found with index in trap-state?	mg  wrapper-at-index?	ng  next-index!?	og  next-ephemeral-index!?	pg  handler-for-index?	qg  ephemeral-handler-for-index?	rg  make-weak-key-hash-table?	sg  *trap-states*?	tg  	hashq-ref?	ug  
hashq-set!?	vg  trap-state-for-vm?	wg  the-vm?	xg  the-trap-state?	yg  set-vm-trace-level!?	zg  map?	{g  and=>?	|g  trap-at-procedure-call?	}g  format?	~f  Breakpoint at ~a?	g  trace-calls-to-procedure? ?g  prefixS? ?f  	Trap ~a: ? ?f  Tracepoint at ~a? ?g  trap-at-source-location? ?f  Breakpoint at ~a:~a? ?g  trap-frame-finish? ?f  Return from ~a? ?f  ~a:~a:~a? ?g  source:file? ?f  unknown file? ?g  source:line-for-user? ?g  source:column? ?f  unknown source location? ?g  source-string? ?g  into?S? ??	?? ?g  instruction?S? ??	?? ??? ? ?g  frame-next-source? ?g  trap-matching-instructions? ?g  frame-address? ?f  Step to different instruction? ?f  #Step to different instruction in ~a? ?f  Step into ~a? ?f  Step out of ~a?C 5       h?;  ?  ]4	
&'65 49 >  "  G   4:i5 ;R;<=        h(   0  ][$  
 6 6     (      g  frame
		# g  idx		# g  	trap-name			# g  default-handler			#  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	1
??		2	??		3	??		4	??		5	??	#	5	?? 		#	  g  nameg  default-trap-handler? C>R4?i@E5FRF h     - 1 3  ? C       g  index
			 g  enabled?			 g  trap				 g  name				 g  defrec-48699c9-8				  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	7
?? 				
	  g  nameg  make-trap-wrapper? CGR4HiFi5IR4Ji4KiFiA54LiFiA55MR4Ji4KiFiB54LiFiB55NR4Ji4KiFiC54LiFiC55OR4Ji4KiFiD54LiFiD55PR4?iQV5WR>W      h?   v  -  1  3  H J (  "  J ?J ?K J (  
"  J ?J ?K J (  	?"  J ?J ?K J (  "  J ?J ?K ? C   n      g  defrec-48699c9-14
	 ? g  _x		( g  _x	6	B g  _x	Q	] g  _x	k	w g  handler	w ? g  next-idx		w ? g  next-ephemeral-idx		w ? g  wrappers		w ?  	g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	=
??	d	A	?? 		 ?


  g  nameg  make-trap-state? CXR4HiWi5YR4Ji4KiWiR54LiWiR55ZR4Ji4KiWiS54LiWiS55[R4Ji4KiWiT54LiWiT55\R4Ji4KiWiU54LiWiU55]RM   h   ?   ]4 545?C      ?       g  t1
		 g  t2		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	C
??		D	??	
	D	??		D	?? 			  g  nameg  trap-wrapper<?? C^RN_`MOc   hP   .  ]	4 5$  4 564 5445 45 >  "  G  45 6     &      g  wrapper
		K g  trap	!	K  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	J
??		K	??		K	??		L	??		L	$??		L	??		M	??	!	M	??	$	N	??	.	N	*??	7	N	??	K	O	?? 		K  g  nameg  enable-trap-wrapper!? CdRNOc_eM       hP   /  ]	4 5$  14 5445 45 >  "  G  45 64 56     '      g  wrapper
		K g  trap		>  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	Q
??		R	??		R	??		T	??		T	??		U	??	!	U	*??	*	U	??	>	V	??	B	S	??	C	S	%??	K	S	?? 		K  g  nameg  disable-trap-wrapper!? CfRc]gM  h0     ]445 44 5 5>  "  G  6
      g  
trap-state
		0 g  wrapper		0  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	X
??		Y	??		Z	??		Z	??		Z	1??		Z	??	"	Y	??	0	[	?? 			0	  g  nameg  add-trap-wrapper!? ChRc]i h    ?   ]45 44 556     ?       g  
trap-state
		 g  wrapper		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	]
??		_	??		_	??		_	??		^	?? 			  g  nameg  remove-trap-wrapper!? CjRN       h   ?   ]4 5$  ?CC    ?       g  wrapper
		 g  level		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	b	??		c	??		c	
??		d	?? 			   C]   h   ?   ]
4 56?       g  
trap-state
		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	a
??		g	??		b	?? 		  g  nameg  trap-state->trace-level? CkR<lM]      hP   k  ]
"  8(  4>  "  G  C4?5?$  ?C?"???4 5"???     c      g  
trap-state
		K g  idx		K g  wrappers			>  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	i
??		j	??		k	??		m	??		m	??		m	??	#	o	??	(	o	 ??	*	o	??	-	o	??	1	k	??	4	p	??	8	r	
??	>	r	??	>	j	??	?	j	??	K	j	?? 		K	  g  nameg  wrapper-at-index? CmR[c  h(   ?   ]	4 5445 ?>  "  G  C?       g  
trap-state
		( g  idx			(  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	t
??		u	??			u	??		v	??		v	+??		v	?? 		(  g  nameg  next-index!? CnR\c      h(     ]	4 5445 ?>  "  G  C?       g  
trap-state
		( g  idx			(  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	y
??		z	??			z	??		{	??		{	5??		{	?? 		(  g  nameg  next-ephemeral-index!? CoRmZMP       h8     ]4L L54L 5$   45456C            g  frame
		2 g  wrapper		2 g  handler			2  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
		??	 ?	??	 ?	??	 ?	??	 ?	??	! ?	??	( ?	??	0 ?	
?? 			2   C  h   ?   ] O C    ?       g  
trap-state
		 g  idx		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	~
?? 			  g  nameg  handler-for-index? CpRmNfj  hP      ]	4LL 5$  ;45$  4>  "  G  "   4L>  "  G  L 6C?       g  frame
		P g  wrapper		P  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	5 ?	??	N ?	?? 
		P   C     h   ?   ] O C  ?       g  
trap-state
		 g  idx		 g  handler			  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
?? 			  g  nameg  ephemeral-handler-for-index? CqR4ri5 sRtsXuv       h@     ]	4 5$  C45 4 >  "  G   6       	      g  vm
		9 g  t		9 g  ts		9  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	9 ?	?? 		9  g  nameg  trap-state-for-vm? CvRvw    h   ?   ] 45 6     ?       g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	??	 ?	?? 		
  g  nameg  the-trap-state? CxRx;ywk  h    ?   ] L$  45 4L 56C      ?       g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?	??	 ?	??	 ?	!??	 ?	*??	 ?	?? 		
   Cyw    h   ?   ] L $  
45 
6C    ?       g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?	??	 ?	??	 ?	!??	 ?	?? 		
   Cywk   h    ?   ] L$  45 4L 56C      ?       g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?	??	 ?	??	 ?	!??	 ?	*??	 ?	?? 		
   Cyw    h   ?   ] L $  
45 
6C    ?       g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?	??	 ?	??	 ?	!??	 ?	?? 		
   C  hp   	  - . , 3 #  45  Y O  O 4 O >   "  G  V4>   X4 O >   "  G  "  ZCZF        g  handler
		n g  thunk		n g  
trap-state			n  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	;??	0 ?	?? 		n	  g  nameg  with-default-trap-handler? C/RxzM]        h(   ?   -  . , 3  #  45  4 56  ?       g  
trap-state
		&  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	,??	 ?	??	& ?	?? 		&
  g  nameg  
list-traps? C)Rx{mP      h(   ?   - . , 3 #  45 4 56?       g  idx
		( g  
trap-state		(  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	/??	 ?		??	( ?	?? 		(  g  nameg  	trap-name? C+Rx{mN       h(   ?   - . , 3 #  45 4 56?       g  idx
		( g  
trap-state		(  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	3??	 ?		??	( ?	?? 		(  g  nameg  trap-enabled?? C*Rx{md   h(   ?   - . , 3 #  45 4 56?       g  idx
		( g  
trap-state		(  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	2??	 ?		??	( ?	?? 		(  g  nameg  enable-trap!? C,Rx{mf    h(   ?   - . , 3 #  45 4 56?       g  idx
		( g  
trap-state		(  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	3??	 ?		??	( ?	?? 		(  g  nameg  disable-trap!? C-Rx{mNfj      h0   ?   ]4 5$  4 >  "  G  "   L  6     ?       g  wrapper
		+  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?		??	 ?	??	 ?	??	 ?	??	+ ?	?? 		+   C   h0   ?   - . , 3 #  45 4 5O 6   ?       g  idx
		- g  
trap-state		-  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	2??	 ?		??	- ?	?? 		-  g  nameg  delete-trap!? C.RxcZ      h(   ?   - . , 3 #  45 45 6  ?       g  handler
		& g  
trap-state		&  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	???	& ?	?? 		&  g  nameg  install-trap-handler!? C0Rxn|phG}~ 	       hP   z  - . , 3 #  45 454 45544 556  r      g  proc
		N g  
trap-state		N g  idx		 	N g  trap		2	N  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	B??	 ?	??	  ?	??	# ?	??	( ?	??	2 ?	??	2 ?	??	9 ?	??	A ?	??	F ?	??	J ?	??	L ?	??	N ?	?? 		N  g  nameg  add-trap-at-procedure-call!? C1Rxn?}?hG? 
    hX   ?  - . , 3 #  45 454 45544	 556       ?      g  proc
		Q g  
trap-state		Q g  idx		 	Q g  trap		5	Q  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	>??	 ?	??	  ?	??	# ?	??	* ?	??	/ ?	$??	3 ?	??	5 ?	??	5 ?	??	< ?	??	D ?	??	I ?	??	M ?	??	O ?	??	Q ?	?? 		Q  g  nameg  add-trace-at-procedure-call!? C2Rxn?phG}? 	   hX   ?  - . , 3 #  45 454 45544 556      ?      g  file
		R g  	user-line		R g  
trap-state			R g  idx		 	R g  trap		4	R  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	>??	 ?	??	  ?	??	# ?	??	* ?	(??	4 ?	??	4 ?	??	; ?	??	C ?	??	H ?	??	N ?	??	P ?	??	R ?	?? 		R	  g  nameg  add-trap-at-source-location!? C3Rxo?q.   h   ?   ]LL 6      ?       g  frame
		
  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?	??	
 ?	 ?? 		
   ChG}? 
       h`   ?  - . , 3 #  45 454 45O 544	 556       ?      g  frame
		Y g  handler		Y g  
trap-state			Y g  idx		 	Y g  trap		=	Y  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
 ?
??	 ?	:??	 ?	??	  ?	??	# ?	??	( ?	??	= ?	??	= ?	??	D ?	??	L	??	Q	??	U	??	W ?	??	Y ?	?? 		Y	  g  nameg  #add-ephemeral-trap-at-frame-finish!? C4R}??????  h@     ]	 $  /4 5$  "  4 54 56C            g  source
		: g  t		'  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?

??		??		??		 ??		??	$	5??	(	??	/	,??	7	??	9	?? 		:  g  nameg  source-string? C?R?x?o?   h   ?   ]C    ?       g  f
		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	?? 		   C?  h   ?   ]4 5L ??C  ?       g  f
		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	??		.??		&??		!?? 		   C??     h   ?   ]4 5L?$  L  6C?       g  f
		  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	
??		??		??		??		?? 		   CqhG?}????       h?   ?  - /   0   3 	#  45 #  #  4 5454$  "  O $  "  4 5	O 4
554$  $  "  
4 5"  ($  4455"  445556 ?      g  frame
	 ? g  handler	 ? g  
trap-state		 ? g  into?		 ? g  instruction?		 ? g  source		9 ? g  idx		B ? g  	predicate		Z ? g  fp		o	} g  trap	 ? ?  
g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
	
??		3??	3	??	9	??	<	??	B	??	E	??	M	??	Z	??	b	??	i	??	o	?? ?	?? ?	?? ?	?? ?	?? ?!	?? ?"	
?? ?#	?? ?$	?? ?$	?? ?$	?? ?%	
?? ?&	?? ?&	?? ?&	(?? ?&	?? ?'	?? ?'	?? ?'	*?? ?'	?? ?	?? ?	?? "	 ?	g  into?S	?g  instruction?S	?   g  nameg  add-ephemeral-stepping-trap!? C5RxnhG      h8   !  - . , 3 #  45 454 56          g  trap
		4 g  name		4 g  
trap-state			4 g  idx		 	4  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?
)
??	)	5??	*	??	 *	??	'-	??	4+	?? 		4	  g  nameg  	add-trap!? C(RC  ?      g  m
		,  g  filenamef  \/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/system/vm/trap-state.scm?		
??	-	/	??	6	/
???	1
???	7
???	=
??u	C
??		J
??
?	Q
??	X
??@	]
??>	a
??	i
??J	t
???	y
???	~
??a ?
??b ?	??k ?
??? ?
??? ?
??? ?
?? ? ?
??" ?
??#Q ?
??$? ?
??%? ?
??'? ?
??), ?
??+ ?
??-( ?
??/6 ?
??2$ ?
??3?
??:&	
??;?)
?? %	;?
   C6 