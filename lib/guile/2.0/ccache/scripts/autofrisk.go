GOOF----LE-4-2.0.      ] | 4  hg
      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  scripts?	g  	autofrisk?	 ?		g  filenameS?	
f  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?	g  importsS?	g  srfi?	g  srfi-1?	 ?	 ?	g  srfi-8?	 ?	 ?	g  srfi-13?	 ?	 ?	g  srfi-14?	 ?	 ?	g  read-scheme-source?	 ?	 ?	g  frisk?	 ?	 ?	 ?	 g  exportsS?	! ?	"g  	autoloadsS?	#g  ice-9?	$g  popen?	%#$ ?	&g  open-input-pipe?	'& ?	(%' ?	)g  set-current-module?	*) ?	+) ?	,g  %include-in-guild-list?	-f  0Generate snippets for use in configure.ac files.?	.g  %summary?	/g  
files-glob?	0g  non-critical-external?	1g  non-critical-internal?	2g  programs?	3g  pww-varname?	4/0123 ?	5g  *recognized-keys*?	6g  error?	7f  syntax error:?	8f  input not a list?	9g  every?	:g  list??	;f  non-list element?	<g  length?	=f  list too short?	>g  quote?	?g  memq?	@f  unrecognized key:?	Ag  apply?	Bg  map?	Cg  fold?	Dg  append?	Eg  assq-ref?	Fg  canonical-configuration?	Gg  for-each?	Hg  format?	If  GUILE_MODULE_REQUIRED~A
?	Jg  >>strong?	Kg  object->string?	Lg  string-map!?	Mg  char-set-contains??	Ng  char-set:letter+digit?	Og  	safe-name?	Pf  probably_wont_work?	Qg  *pww*?	Rg  edge-up?	Sg  	edge-down?	Tf  have_guile_module~A?	Uf  GUILE_MODULE_AVAILABLE(~A, ~A)
?	Vf  "test "$~A" = no &&
  ~A="~A $~A"~A?	Wf  

?	Xg  >>weak?	Yf  guile_module~Asupport_~A?	Zf  AC_PATH_PROG(~A, ~A)
?	[f  test \
?	\f   "$~A" = "" -o \
?	]f  ~A &&
~A="~A $~A"

?	^g  list-ref?	_f  war = peace?	`f  freedom = slavery?	af  ignorance = strength?	bg  random?	cg  	>>program?	dg  
>>programs?	ef  echo '(' ~A ')'?	fg  symbol->string?	gg  read?	hg  unglob?	ig  make-frisker?	jg  external?	kg  	partition?	lg  member?	mg  mod-down-ls?	nf  AC_DEFUN([AUTOFRISK_CHECKS],[

?	of  
~A=~S

?	pf   ?	qf  AC_SUBST(~A)
])

?	rg  >>checks?	sg [
AC_DEFUN([AUTOFRISK_SUMMARY],[
if test ! "$~A" = "" ; then
    p="         ***"
    echo "$p"
    echo "$p NOTE:"
    echo "$p The following modules probably won't work:"
    echo "$p   $~A"
    echo "$p They can be installed anyway, and will work if their"
    echo "$p dependencies are installed later.  Please see README."
    echo "$p"
fi
])
?	tg  	>>summary?	uf  
modules.af?	vg  file-exists??	wf  could not find input file:?	xg  with-output-to-file?	yf  ~A.m4?	zg  read-scheme-source-silently?	{g  main?C 5    h?!  2  ]4	
 !"(5	 4+ >  "  G   ,R-.R45R6789:;<       h   ?   ]4 5?C    ?       g  form
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	N	??		N	$??		N	?? 		   C=>?5     h8     ] ??$  $&  C45$  CN CC          g  form
		4 g  key		4 g  t		!	2  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	P	??		Q	??		Q	??	
	R	??		R	??		S	$??		S	??		S	??		T	??	!	S	??	0	V	!?? 		4   C@ABCD   h8     ]
 ?L &   ?$  4 ?5"  "  $  CC    ?       g  form
		4 g  so-far		4 g  t		&	4  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	[	 ??		\	0??		\	&??		]	+??		\	&??		^	+??		^	:??		^	+??	&	\	"?? 
		4	   C   h   ?   ] O   L 6    ?       g  key
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	Z	??		`	 ??		[	?? 		   C5E  h   ?   ]L  6      ?       g  key
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	c	??	
	d	?? 		
   C       h?   9  ] ?$  "  4>  "  G  4 5$  "  4>  "  G  4 5$  "  4>  "  G  H4	O  5
J $  "  4>  "  G  4 O 5O C  1      g  forms
	 ? g  	condition	*	M g  	condition	V	y g  un	z ? g  	condition	 ? ? g  x	 ? ? g  bunched ? ?  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	I
??		L		??			K	??		K	??		K	*??		L	??		K	??	"	M		??	*	M	??	2	K	??	7	K	??	;	K	*??	=	M	??	B	K	??	N	N		??	V	N	??	^	K	??	c	K	??	g	K	*??	i	N	;??	n	K	??	z	O	??	}	P	?? ?	Y	?? ?	P	?? ?	K	?? ?	K	?? ?	K	*?? ?	K	?? ?	Z	?? ?	Z	??  	 ?  g  nameg  canonical-configuration? CFRGHI h   ?   ] 6     ?       g  module
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	g	??		h	??		h	?? 		   C    h   ?   ] 6      ?       g  modules
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	f
??	
	g	?? 		
  g  nameg  >>strong? CJRKLMN  h   ?   ]4 5$   C_C   ?       g  c
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	m	??		n	??		n	?? 		   C h(   ?   ]	4 54>  "  G  C      ?       g  module
		" g  var			"  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	k
??		l	??			l	??		m	?? 		"  g  nameg  	safe-name? CORPQRGRSHTOUVQW 
hP   I  ]4 54 544554>  "  G  	6   A      g  edge
		M g  up			M g  down			M g  var		#	M  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	w	??		x	??			x	??		y	??		x	??		z	??		z	%??		z	;??	#	z	??	#	x	??	&	{	??	+	{	??	4	{	??	A	|	??	K	}	-??	M	|	?? 		M   C    h   ?   ] 6      ?       g  
weak-edges
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
	v
??	
	w	?? 		
  g  nameg  >>weak? CXRBHYO h   ?   ]4L 5 6      ?       g  prog
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	 ??	 ?	??	 ?	?? 		   CGHZ   h   ?   ] 6   ?       g  var
		 g  prog		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	??	 ?	?? 			   CH[H\ h   ?   ] 6     ?       g  var
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	??	 ?	?? 		   C]^_`abQ      hx   l  ]
4 O 54>  "  G  4>  "  G  4>  "  G  4	
 4	55 6       d      g  module
		q g  progs		q g  vars			q  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?
??	 ?	??	 ?	??	 ?	??	) ?	??	. ?	??	3 ?	??	< ?	??	T ?	??	U ?	??	Y ?	??	[ ?	??	] ?	??	` ?	??	a ?	??	i ?	??	q ?	?? 		q	  g  nameg  	>>program? CcRGc        h   ?   ] ? ?6    ?       g  form
		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	??	
 ?	$??	 ?	?? 		   C  h   ?   ] 6      ?       g  programs
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?
??	
 ?	?? 		
  g  nameg  
>>programs? CdR&HeBfg  h      ]	44 55456      g  pattern
		  g  p		   g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?
??	 ?	??	 ?	??	 ?	'??	 ?	??	 ?	??	 ?	??	 ?	??	  ?	?? 
		   g  nameg  unglob? ChRFDBh/01ij3Qkl9l      h   ?   ] L 6      ?       g  i
		
  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	 ??	
 ?	"?? 		
   CBSm h0   ?   ]	4 L 5$  CLO 44 556 ?       g  module
		/ g  t		/  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	??	 ?	??	  ?	 ??	% ?	/??	- ?	 ??	/ ?	?? 		/   CHnJopXCDm    h   ?   ]4 56 ?       g  module
		 g  so-far		  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	%??	 ?	?? 			   Cd2q       h?   o  ]A4 544455?4545445 54	54
5(  "  ? 4O >  G 4>  "  G  4>  "  G  4>  "  G  44  5>  "  G  445>  "  G  6   g      g  forms
	 ? g  cfg		 ? g  files		 ? g  ncx		' ? g  nci		0 ? g  report		< ? g  external		E ? g  pww-varname		N	b g  weak		v ? g  strong		v ?  
g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	*??	 ?	/??	 ?	*??	 ?	??	 ?	??	 ?	??	! ?	??	% ?	??	' ?	??	' ?	??	* ?	??	. ?	??	0 ?	??	0 ?	??	3 ?	??	4 ?	??	< ?	??	< ?	??	? ?	??	C ?	??	E ?	??	E ?	??	H ?	??	L ?	??	N ?	??	N ?	??	V ?	??	] ?	*??	_ ?	??	c ?	??	y ?	??	~ ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	%?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? 9	 ?  g  nameg  >>checks? CrRHfsQ  h   ?   ] 456      ?       g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?
??	 ?	
??	
 ?	??	 ?	
??	 ?	?? 		
  g  nameg  	>>summary? CtRuv6wxHyrzt      h    ?   ] 44L 5>  "  G  6    ?       g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?	??	 ?	??	 ?	??	 ?	??	 ?	?? 		
   C 	   h`   L  -  1  3  (  "   ?45$  "  4>  "  G  45O 6       D      g  args
			Y g  file		Y g  t		!	D  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?
 ?
??	 ?	??	 ?	??	 ?	,??	 ?	??	 ?	??	! ?	??	. ?	??	2 ?	??	9 ?	??	G ?	??	L ?	$??	P ?	??	Y ?	?? 			Y


  g  nameg  	autofrisk? CRi{RC*      g  m
		0  g  filenamef  Y/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/autofrisk.scm?		6
??	4	@
??	6	A	??	9	A
??	;	C	??	>	C
???	I
??
I	f
??9	k
??;	t	??>	t
???	v
??? ?
??N ?
??? ?
??" ?
?? ?
??!? ?
??!? ?
?? 	!?
   C6 