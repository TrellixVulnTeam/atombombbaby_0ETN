GOOF----LE-4-2.0(5      ] ' 4       ho      ] g  apply:nconc2last?	g  apply?	g  call-with-current-continuation?	g  call-with-values?	g  dynamic-wind?	f  r?	g  	OPEN_READ?	f  w?		g  
OPEN_WRITE?	
f  r+?	g  	OPEN_BOTH?	f  	/dev/null?	g  *null-device*?	g  	open-file?	g  open-input-file?	g  open-output-file?	g  open-io-file?	g  close-input-port?	g  call-with-input-file?	g  close-output-port?	g  call-with-output-file?	g  set-current-input-port?	g  with-input-from-port?	g  set-current-output-port?	g  with-output-to-port?	g  set-current-error-port?	g  with-error-to-port?	g  with-input-from-file?	g  with-output-to-file?	g  with-error-to-file?	g  call-with-input-string?	 g  with-input-from-string?	!g  call-with-output-string?	"g  with-output-to-string?	#g  with-error-to-string?	$f   ?	%g  	read-char?	&g  the-eof-object?C 5    h?/  ?  ]         h   ?   - 1 3  45@    ?       g  fun
			 g  args			  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
	!
??		"	??		"	?? 			
  g  nameg  apply? CRh   ?   ] B   ?       g  proc
		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
	#
??		$	?? 		  g  nameg  call-with-current-continuation? CR h   ?   ]4 >   6< ?       g  producer
		 g  consumer		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
	%
??		&	?? 			  g  nameg  call-with-values? CR     hH   ?  ] 4 >   "  G  V4>   X4>   "  G  CX4>   "  G  F     ?      g  in
		C g  thunk		C g  out			C  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
	'
??		Y	??		Y	??	!	Y	?? 		C	  g  nameg  dynamic-wind?g  documentationf ?All three arguments must be 0-argument procedures.
Guard @var{in} is called, then @var{thunk}, then
guard @var{out}.

If, any time during the execution of @var{thunk}, the
continuation of the @code{dynamic_wind} expression is escaped
non-locally, @var{out} is called.  If the continuation of
the dynamic-wind is re-entered, @var{in} is called.  Thus
@var{in} and @var{out} may be called any number of
times.
@lisp
 (define x 'normal-binding)
@result{} x
 (define a-cont
   (call-with-current-continuation
     (lambda (escape)
       (let ((old-x x))
         (dynamic-wind
           ;; in-guard:
           ;;
           (lambda () (set! x 'special-binding))

           ;; thunk
           ;;
           (lambda () (display x) (newline)
                   (call-with-current-continuation escape)
                   (display x) (newline)
                   x)

           ;; out-guard:
           ;;
           (lambda () (set! x old-x)))))))

;; Prints:
special-binding
;; Evaluates to:
@result{} a-cont
x
@result{} normal-binding
 (a-cont #f)
;; Prints:
special-binding
;; Evaluates to:
@result{} a-cont  ;; the value of the (define a-cont...)
x
@result{} normal-binding
a-cont
@result{} special-binding
@end lisp? CRR	R
RR        h   v  ] 6      n      g  str
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
	t
??	
	x	?? 		
  g  nameg  open-input-file?g  documentationf  ?Takes a string naming an existing file and returns an input port
capable of delivering characters from the file.  If the file
cannot be opened, an error is signalled.? CR	       h   ?  ] 6      ?      g  str
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
	z
??	
		?? 		
  g  nameg  open-output-file?g  documentationf Takes a string naming an output file to be created and returns an
output port capable of writing characters to a new file by that
name.  If the file cannot be opened, an error is signalled.  If a
file with the given name already exists, the effect is unspecified.? CR     h     ] 6      ?       g  str
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	
 ?	?? 		
  g  nameg  open-io-file?g  documentationf  2Open file with name STR for both input and output.? CR    h8   |  ]4 54>  G4>  "  G  E       t      g  str
		1 g  proc		1 g  p				1 g  vals			1  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	1 ?	?? 		1	  g  nameg  call-with-input-file?g  documentationf EPROC should be a procedure of one argument, and STR should be a
string naming a file.  The file must
already exist. These procedures call PROC
with one argument: the port obtained by opening the named file for
input or output.  If the file cannot be opened, an error is
signalled.  If the procedure returns, then the port is closed
automatically and the values yielded by the procedure are returned.
If the procedure does not return, then the port will not be closed
automatically unless it is possible to prove that the port will
never again be used for a read or write operation.? CR h8   ?  ]4 54>  G4>  "  G  E       ?      g  str
		1 g  proc		1 g  p				1 g  vals			1  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	??		 ?	??	 ?	??	 ?	??	 ?	??	1 ?	?? 		1	  g  nameg  call-with-output-file?g  documentationf bPROC should be a procedure of one argument, and STR should be a
string naming a file.  The behaviour is unspecified if the file 
already exists. These procedures call PROC
with one argument: the port obtained by opening the named file for
input or output.  If the file cannot be opened, an error is
signalled.  If the procedure returns, then the port is closed
automatically and the values yielded by the procedure are returned.
If the procedure does not return, then the port will not be closed
automatically unless it is possible to prove that the port will
never again be used for a read or write operation.? CR    h   ?   ] 4M 5N C   ?       g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	 ?	)??	 ?	?? 		
  g  nameg  swaports? C        hX   ?   ]
 H O  Q 4>   "  G  V4>   X4>   "  G  CX4>   "  G  F    ?       g  port
		T g  thunk		T g  swaports			T  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	??	 ?	?? 		T	  g  nameg  with-input-from-port? CR  h   ?   ] 4M 5N C   ?       g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	 ?	)??	 ?	?? 		
  g  nameg  swaports? C        hX   ?   ]
 H O  Q 4>   "  G  V4>   X4>   "  G  CX4>   "  G  F    ?       g  port
		T g  thunk		T g  swaports			T  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	??	 ?	?? 		T	  g  nameg  with-output-to-port? CR   h   ?   ] 4M 5N C   ?       g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	 ?	)??	 ?	?? 		
  g  nameg  swaports? C        hX   ?   ]
 H O  Q 4>   "  G  V4>   X4>   "  G  CX4>   "  G  F    ?       g  port
		T g  thunk		T g  swaports			T  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	??	 ?	?? 		T	  g  nameg  with-error-to-port? CR  h   ?   ] L 6      ?       g  p
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	
 ?	?? 		
   C      h   ?  ] O 6 ?      g  file
		 g  thunk		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	?? 			  g  nameg  with-input-from-file?g  documentationf THUNK must be a procedure of no arguments, and FILE must be a
string naming a file.  The file must already exist. The file is opened for
input, an input port connected to it is made
the default value returned by `current-input-port', 
and the THUNK is called with no arguments.
When the THUNK returns, the port is closed and the previous
default is restored.  Returns the values yielded by THUNK.  If an
escape procedure is used to escape from the continuation of these
procedures, their behavior is implementation dependent.? CR    h   ?   ] L 6      ?       g  p
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	
 ?	?? 		
   C      h     ] O 6       g  file
		 g  thunk		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	?? 			  g  nameg  with-output-to-file?g  documentationf *THUNK must be a procedure of no arguments, and FILE must be a
string naming a file.  The effect is unspecified if the file already exists. 
The file is opened for output, an output port connected to it is made
the default value returned by `current-output-port', 
and the THUNK is called with no arguments.
When the THUNK returns, the port is closed and the previous
default is restored.  Returns the values yielded by THUNK.  If an
escape procedure is used to escape from the continuation of these
procedures, their behavior is implementation dependent.? CR        h   ?   ] L 6      ?       g  p
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	
 ?	?? 		
   C      h     ] O 6 
      g  file
		 g  thunk		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	?? 			  g  nameg  with-error-to-file?g  documentationf )THUNK must be a procedure of no arguments, and FILE must be a
string naming a file.  The effect is unspecified if the file already exists. 
The file is opened for output, an output port connected to it is made
the default value returned by `current-error-port', 
and the THUNK is called with no arguments.
When the THUNK returns, the port is closed and the previous
default is restored.  Returns the values yielded by THUNK.  If an
escape procedure is used to escape from the continuation of these
procedures, their behavior is implementation dependent.? CR  h   ?   ] L 6      ?       g  p
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	
 ?	?? 		
   C      h   g  ] O 6 _      g  string
		 g  thunk		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	?? 			  g  nameg  with-input-from-string?g  documentationf xTHUNK must be a procedure of no arguments.
The test of STRING  is opened for
input, an input port connected to it is made, 
and the THUNK is called with no arguments.
When the THUNK returns, the port is closed.
Returns the values yielded by THUNK.  If an
escape procedure is used to escape from the continuation of these
procedures, their behavior is implementation dependent.? C R!     h   ?   ] L 6      ?       g  p
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	
 ?	?? 		
   C      h   	  ] O 6         g  thunk
		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	?? 		  g  nameg  with-output-to-string?g  documentationf  /Calls THUNK and returns its output as a string.? C"R!   h   ?   ] L 6      ?       g  p
		
  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	??	
 ?	?? 		
   C      h     ] O 6         g  thunk
		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?
??	 ?	?? 		  g  nameg  with-error-to-string?g  documentationf  5Calls THUNK and returns its error output as a string.? C#R4i$%  h   ?   ] 6?       g  p
		  g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm?
 ?	2??	 ?	>?? 		   C5&RC  ~      g  filenamef  R/pub/devel/guile/yaakov/guile-2.0.14-1.i686/src/guile-2.0.14/module/ice-9/r4rs.scm? ?	!
???	#
???	%
???	'
???	n	???	n
???	o	???	o
???	p	???	p
???	r	??	 	r
??
?	t
???	z
??? ?
??? ?
??} ?
??? ?
??? ?
?? ?
??? ?
??#? ?
??'? ?
??+2 ?
??-$ ?
??/ ?
??/ ?	??/  ?	/??/? ?	??/? ?
?? 	/?
   C6 