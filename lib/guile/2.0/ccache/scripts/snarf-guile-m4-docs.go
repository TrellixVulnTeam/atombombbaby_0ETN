GOOF----LE-4-2.0      ] 3 4   h?      ] g  guile?	 ?	g  define-module*?	 ?	 ?	g  scripts?	g  snarf-guile-m4-docs?	 ?		g  filenameS?	
f  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?	g  importsS?	g  ice-9?	g  rdelim?	 ?	 ?	 ?	g  exportsS?	 ?	g  set-current-module?	 ?	 ?	g  %include-in-guild-list?	f  /Snarf out texinfo documentation from .m4 files.?	g  %summary?	g  display?	f  @deffn {Autoconf Macro}?	g  for-each?	g  string=??	f  #?	g  	substring?	g  string-length?	 f  # ?	!g  newline?	"f  
@end deffn?	#g  display-texi?	$g  catch?	%$ ?	&$ ?	'g  prefix??	(g  list->string?	)g  reverse?	*g  string->list?	+g  massage-usage?	,g  	open-file?	-f  r?	.g  eof-object??	/f  # Usage:?	0g  	read-line?	1f  AC_DEFUN?	2g  main?C 5    h?  ?   ]4	
5 4 >  "  G   RR !  h?   A  ]4"  &44 
55$  4 5"  B "  <4 5	?$  *44 
	55$  4 	5"  "???"  "???>  "  G  6     9      g  line
		|  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	*	??		+	??	
	.	??		.	(??		.	,??		.	??		+	??		/	??	/	+	??	0	+	'??	9	+	#??	=	+	??	>	,	#??	B	,	-??	C	,	2??	N	,	#??	R	+	??	S	-	??	p	+	??	|	1	?? 		|   C"!        hP   
  ]4>  "  G  4 >  "  G  4>  "  G  4>   "  G  6         g  lines
		N  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	(
??		)	??		)	??		)	??		*	??	)	3	??	-	3	??	2	3	??	;	4	??	N	4	?? 		N  g  nameg  display-texi? C#R&      h   ?   ] L4L 
4L556 ?       g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	7	??		8	??		8	#??		8	??		8	?? 		
   Ch   ?   -  1  3 C     ?       g  args
			  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	7	?? 			


   C  h   ?   ] O 6      ?       g  line
		 g  sub		  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	6
??		7	?? 			  g  nameg  prefix?? C'R()*      hp   u  ]"  Y(  4455 C??(?$  "  )?$  "  ,?$   "  ?"???4 5"??? m      g  line
		o g  line		_ g  acc			_ g  key		"	T  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	:
??		;	??		<	??		=	??		=	??		=	??		=	??		>	??	"	?	??	"	?	??	Q	A	??	W	?	??	_	>	??	_	;	??	`	;	??	g	;	-??	o	;	?? 		o  g  nameg  massage-usage? C+R,-.'/0+1#)    h?   V  -  1  3 4 ?5"  ?45$  C45$  4544	55"???4	5$  '4
45>  "  G  45"???"  45"??|$  %45$  45?"??U"???"???45"??=    N      g  args
		 ? g  p	 ? g  line		 ? g  acc		 ? g  t		  ?  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?
	D
??	
	E	??		E	??		E	"??		E	??		E	??		G	??		H	
??	 	H	??	,	I	??	2	I	??	4	I	??	8	I	
??	9	F	??	@	J	??	C	J	-??	M	J	??	U	J	??	V	K	??	\	K	??	^	K	??	b	I	
??	c	L	??	f	L	??	q	L	??	z	F	?? ?	M	?? ?	F	?? ?	Q	?? ?	I	
?? ?	N	?? ?	N	(?? ?	N	?? ?	N	?? ?	F	?? ?	O	?? ?	O	?? ?	G	?? ?	F	?? ?	G	?? )		 ?


  g  nameg  snarf-guile-m4-docs? CRi2RC      ?       g  m
		,  g  filenamef  c/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/scripts/snarf-guile-m4-docs.scm?		!
??	0	%
??	2	&	??	5	&
???	(
??8	6
??8	:
???	D
???	S
?? 
	?
   C6 