GOOF----LE-4-2.0?&      ] p 4      h?	      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  scripts?	g  help?	 ?		g  filenameS?	
f  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?	g  importsS?	g  ice-9?	g  format?	 ?	 ?	g  documentation?	 ?	 ?	g  srfi?	g  srfi-1?	 ?	g  selectS?	g  fold?	g  
append-map?	 ?	 ?	 ?	g  exportsS?	g  	show-help?	g  show-summary?	g  
show-usage?	 g  main?	!  ?	"g  set-current-module?	#" ?	$" ?	%f  Show a brief help message.?	&g  %summary?	'f  help
help --all
help COMMAND?	(g  	%synopsis?	)f  ?
Show help on guild commands.  With --all, show arcane incantations as
well.  With COMMAND, show more detailed help for a particular command.
?	*g  %help?	+g  file-exists??	,g  file-is-directory??	-g  opendir?	.g  eof-object??	/g  closedir?	0g  readdir?	1g  string=??	2f  .?	3f  ..?	4g  directory-files?	5g  or-map?	6g  string-suffix??	7g  string-null??	8g  	substring?	9g  string-length?	:g  append?	;g  %load-compiled-extensions?	<g  %load-extensions?	=g  strip-extensions?	>g  unique?	?g  map?	@g  symbol->string?	Ag  sort?	Bg  in-vicinity?	Cg  
%load-path?	Dg  string<??	Eg  find-submodules?	Fg  display?	Gf  lUsage: guild COMMAND [ARGS]
Run command-line scripts provided by GNU Guile and related programs.

Commands:
?	Hg  for-each?	Ig  string->symbol?	Jg  resolve-module?	Kg  ensureS?	Lg  and=>?	Mg  module-variable?	Ng  variable-ref?	Og  %include-in-guild-list?	Pf    ~A ~23t~a
?	Qf    ~A
?	R ?	Sf 
For help on a specific command, try "guild help COMMAND".

Report guild bugs to ~a
GNU Guile home page: <http://www.gnu.org/software/guile/>
General help using GNU software: <http://www.gnu.org/gethelp/>
For complete documentation, run: info guile 'Using Guile Tools'
?	Tg  %guile-bug-report-address?	Ug  list-commands?	Vg  file-commentary?	Wg  %search-load-path?	Xg  module-filename?	Yg  module-commentary?	Zg  	last-pair?	[g  module-name?	\g  module-command-name?	]g  current-output-port?	^g  string-split?	_g  string-append?	`f  
 OPTION...?	af  Usage: guild ?	bg  newline?	cf         guild ?	df  )No documentation found for command "~a".
?	eg  current-module?	fg  %mod?	gf  --all?	hg ?	if  -a?	ji ?	kg  current-error-port?	lg  exit?	mg  string-prefix??	nf  -?	of  No command named "~a".
?C 5       h  3  ]4	
!5 4$ >  "  G   %&R'(R)*R+,-./0123 
      h?   ?  ]!4 5$  ?4 5$  }4 5"  `45$  4>  "  G  C4545$  "  	4	5$  "  ?"???45"???CC       ?      g  dir
	 ? g  
dir-stream	 ? g  new		% ? g  acc		% ? g  t		U	n  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	)
??		*	??		*	??		*	??		*	??		+	??		+	??	%	,	??	&	.	??	0	.	
??	1	0	??	F	2	??	M	3	??	Q	3	&??	U	3	??	U	3	??	c	4	??	g	4	&??	k	4	??	r	3	??	}	6	?? ?	2	?? ?	,	?? ?	,	?? ?	-	?? ?	,	?? ?	7	?? 	 ?  g  nameg  directory-files? C4R56789       h8   ?   ]4 L 5$  #4 5$  CL 
4L 54 5?6C    ?       g  ext
		4  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	:	
??		<	??		;	??		@	??		;	??	"	B	??	)	B	0??	0	B	??	2	A	?? 
		4   C:;<      h   ?   ] O 456  ?       g  path
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	9
??		C	
??		:	?? 		  g  nameg  strip-extensions? C=R> h8     ] (   C ?(   C ? ???$   ?6 ?4 ?5?C         g  l
		5  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	E
??		F	??		G	??		F	??		H	??		H	??		H		??	!	F	??	&	H	+??	(	H	#??	+	I	??	,	I	??	1	I	$??	3	I	??	4	I	?? 		5  g  nameg  unique? C>R?@>A=        h    ?   ]
4 5$  ?CC      ?       g  x
		 g  rest		 g  stripped				  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	P	??		Q	,??			Q	??		R	??		R	+?? 			   C4B h   ?   ] 6      ?       g  x
		
 g  y		
  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	U	!??	
	U	/?? 		
	   C       h    ?   ]44 L 556       ?       g  path
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	O	??		S	??		T	??		U	??		T	??		P	?? 		   CCD 	      h(   ?   ]	4 544O 556  ?       g  head
		& g  shead		&  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	K
??		L	??		L	??		N	??		O	??	$	N	??	&	M	?? 		&  g  nameg  find-submodules? CERFGHIJKLM&NOPQ   h?   ?  ]!4 5 45$  4455"  $  GL $  "  4	5$  "  $  $  
 6
 6CC?      g  name
	 ? g  modname	 ? g  mod		 ? g  summary		7 ? g  v		S	f  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	b	??		c	??		c	 ??		c	??		c	??		d	??		c	??	"	e	??	#	e	??	&	e	%??	,	e	:??	.	e	%??	2	e	??	7	c	??	?	g	??	E	h	??	K	i	??	Q	i	2??	S	i	??	S	i	??	[	j	??	^	j	??	j	g	??	p	k	??	u	l	??	{	l	?? ?	m	?? ?	m	?? 	 ?   CERST 
h@     ]4>  "  G  4 O 45>  "  G  	6     ?       g  all?
		;  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	Y
??		Z	??		Z	??		Z	??		a	??		n	??	#	n	??	%	n	??	*	a	??	7	o	??	;	o	?? 		;  g  nameg  list-commands? CURVWX        h   ?   ]44 556      ?       g  mod
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	x
??		z	??		z	??		z	??		y	?? 		  g  nameg  module-commentary? CYR@Z[     h   ?   ]44 55?6     ?       g  mod
		  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	|
??		}	??		}	"??		}	??		}	??		}	?? 		  g  nameg  module-command-name? C\R]^M(_\`FabHFcb    h0   ?   ]4L >  "  G  4 L >  "  G  L 6?       g  u
		0  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	0 ?	?? 		0   C        h?   ?  - . , 3 #  45 44 5$  "  44 55
54	>  "  G  4?>  "  G  4
>  "  G  O ?6      ?      g  mod
	 ? g  port	 ? g  var		%	E g  usages		I ?  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
	
??			*??	 ?	??	 ?	??	# ?	1??	% ?	??	% ?	??	- ?	??	0 ?	??	5 ?	??	8 ?	&??	@ ?	&??	B ?	??	I ?	??	I ?	??	L ?	??	P ?	??	W ?	??	` ?	??	e ?	??	j ?	??	s ?	?? ? ?	?? ? ?	?? 	 ?  g  nameg  
show-usage? CR]M&Fb     hH   /  - . , 3 #  45 4 5$  4>  "  G  6C '      g  mod
		G g  port		G g  var		"	G  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
 ?
??	 ?	,??	 ?	??	  ?	"??	" ?	??	" ?	??	* ?	??	+ ?	
??	0 ?	??	7 ?	
??	E ?	
?? 		G  g  nameg  show-summary? CR]M*FbYd\    h?   ?  - . , 3 #  45 4 >  "  G  4 >  "  G  4 5$  4>  "  G  64 5$  4>  "  G  6	
4 56    ?      g  mod
	 ? g  port	 ? g  t		J ? g  t		t ?  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
 ?
??	 ?	)??	 ?	??	. ?	??	B ?	??	H ?	??	J ?	??	J ?	??	S ?		??	X ?	??	_ ?		??	m ?		??	n ?	??	t ?	??	} ?		?? ? ?		?? ? ?	?? ? ?	?? ? ?	?? 	 ?  g  nameg  	show-help? CR4ei5 fRUhjfklmnJIKo       h?     -  1  3  (  6 ?$  "   ?$  6"  445 >  "  G  6 ?(  d4	 ?5$  "??? ?4
45 5$  4>  "  G  
64>  "  G  6"??u       g  args
		 ? g  name	l ? g  t	 ? ?  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?
 ?
??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	 ?	??	& ?	.??	' ?	!??	+ ?	??	0 ?	??	5 ?	??	: ?	??	C ?	??	P ?	??	P ?	??	S ?	??	W ?	??	X ?	!??	\ ?	1??	_ ?	5??	a ?	!??	e ?	??	l ?	??	l ?	??	o ?	??	s ?	??	t ?	#??	} ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? ? ?	?? &		 ?


  g  nameg  main? C RC     +      g  m
		,  g  filenamef  T/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/help.scm?		
??	.	!	??	1	!
??	3	"	??	6	"
??	8	#	??	;	#
???	)
??	9
??d	E
??
?	K
??V	Y
??a	x
??s	|
???	
??? ?
??? ?
??? ?	??? ?
??	 ?
?? 	
   C6 