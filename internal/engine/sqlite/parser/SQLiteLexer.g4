/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2020 by Martin Mirchev
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
 * associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute,
 * sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or
 * substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT
 * NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 *
 * Project : sqlite-parser; an ANTLR4 grammar for SQLite https://github.com/bkiers/sqlite-parser
 * Developed by : Bart Kiers, bart@big-o.nl
 */

// $antlr-format alignTrailingComments on, columnLimit 150, maxEmptyLinesToKeep 1, reflowComments off, useTab off
// $antlr-format allowShortRulesOnASingleLine on, alignSemicolons ownLine

lexer grammar SQLiteLexer;

SCOL:      ';';
DOT:       '.';
OPEN_PAR:  '(';
CLOSE_PAR: ')';
COMMA:     ',';
ASSIGN:    '=';
STAR:      '*';
PLUS:      '+';
MINUS:     '-';
TILDE:     '~';
PIPE2:     '||';
DIV:       '/';
MOD:       '%';
LT2:       '<<';
GT2:       '>>';
AMP:       '&';
PIPE:      '|';
LT:        '<';
LT_EQ:     '<=';
GT:        '>';
GT_EQ:     '>=';
EQ:        '==';
NOT_EQ1:   '!=';
NOT_EQ2:   '<>';

// http://www.sqlite.org/lang_keywords.html
ABORT_:             A B O R T;
ACTION_:            A C T I O N;
ADD_:               A D D;
AFTER_:             A F T E R;
ALL_:               A L L;
ALTER_:             A L T E R;
ANALYZE_:           A N A L Y Z E;
AND_:               A N D;
AS_:                A S;
ASC_:               A S C;
ATTACH_:            A T T A C H;
AUTOINCREMENT_:     A U T O I N C R E M E N T;
BEFORE_:            B E F O R E;
BEGIN_:             B E G I N;
BETWEEN_:           B E T W E E N;
BY_:                B Y;
CASCADE_:           C A S C A D E;
CASE_:              C A S E;
CAST_:              C A S T;
CHECK_:             C H E C K;
COLLATE_:           C O L L A T E;
COLUMN_:            C O L U M N;
COMMIT_:            C O M M I T;
CONFLICT_:          C O N F L I C T;
CONSTRAINT_:        C O N S T R A I N T;
CREATE_:            C R E A T E;
CROSS_:             C R O S S;
CURRENT_DATE_:      C U R R E N T '_' D A T E;
CURRENT_TIME_:      C U R R E N T '_' T I M E;
CURRENT_TIMESTAMP_: C U R R E N T '_' T I M E S T A M P;
DATABASE_:          D A T A B A S E;
DEFAULT_:           D E F A U L T;
DEFERRABLE_:        D E F E R R A B L E;
DEFERRED_:          D E F E R R E D;
DELETE_:            D E L E T E;
DESC_:              D E S C;
DETACH_:            D E T A C H;
DISTINCT_:          D I S T I N C T;
DROP_:              D R O P;
EACH_:              E A C H;
ELSE_:              E L S E;
END_:               E N D;
ESCAPE_:            E S C A P E;
EXCEPT_:            E X C E P T;
EXCLUSIVE_:         E X C L U S I V E;
EXISTS_:            E X I S T S;
EXPLAIN_:           E X P L A I N;
FAIL_:              F A I L;
FOR_:               F O R;
FOREIGN_:           F O R E I G N;
FROM_:              F R O M;
FULL_:              F U L L;
GLOB_:              G L O B;
GROUP_:             G R O U P;
HAVING_:            H A V I N G;
IF_:                I F;
IGNORE_:            I G N O R E;
IMMEDIATE_:         I M M E D I A T E;
IN_:                I N;
INDEX_:             I N D E X;
INDEXED_:           I N D E X E D;
INITIALLY_:         I N I T I A L L Y;
INNER_:             I N N E R;
INSERT_:            I N S E R T;
INSTEAD_:           I N S T E A D;
INTERSECT_:         I N T E R S E C T;
INTO_:              I N T O;
