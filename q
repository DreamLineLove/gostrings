[33mcommit 5030385c54686fb87142e32937a9f31972a93a3c[m[33m ([m[1;36mHEAD -> [m[1;32mmain[m[33m, [m[1;31morigin/main[m[33m)[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 12:08:32 2023 +0630

    Explored EqualFold
    Apparently should not be used in conjunction with other functions like Contains because
      1. Contains takes strings as arguments (and returns a bool)
      2. EqualFold does not return strings. It just checks 2 strings case-insensitively and returns the result as a bool.
    If I want to use contains case-insensitively, I should use ToLower or ToUpper on its arguments.
      Idea: Maybe if I need this contains a lot, I can wrap Contains with a custom function that ToLowers or ToUppers its argumennts, passes them into Contains
      and then returns its result as a bool (from the wrapper function).

[33mcommit b84e05b48bd752b26ef7c5b03b69157e3ae72c17[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 11:45:13 2023 +0630

    Put Cut functions into its own file

[33mcommit 76ab4051448d33592b62684232f2bfd8efc46141[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 11:42:01 2023 +0630

    Explored Cut, CutPrefix, and CutSuffix

[33mcommit 3c4faa73dda5ca6250e2f643e06bd4ff5040a5d3[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 10:07:20 2023 +0630

    Explored Count function

[33mcommit 24d622df707b8cf95746269fcc0ef681c619aba2[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 09:51:02 2023 +0630

    Removed redundant file

[33mcommit a993aafaeb39d19db27c0194ea518eb923796be5[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 09:49:04 2023 +0630

    Explored ToLower and ToUpper

[33mcommit b771c7b62913d112ded83c7ab62a95f12102f53c[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Mon Jun 5 09:42:45 2023 +0630

    Explored the three Contains functions. Also explored ToLower and ToUpper

[33mcommit d94f8628d177360a7f1d5bfbfea0c6d11081f73f[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sun Jun 4 16:31:39 2023 +0630

    Tried out Contains, ToUpper, ToLower, and ContainsAny

[33mcommit 076372a87e128fce98cfae298b44cc8f99e587dd[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 23:15:12 2023 +0630

    Included a helpful readme file with ideas for future development

[33mcommit dc4beba190dcfa0f14048cc8779253cda0c6910f[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 23:10:42 2023 +0630

    Removed import not used error from explore.go

[33mcommit 8da0cf599063a90ac01a7b3f5518c4920ebc4c35[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 23:08:39 2023 +0630

    Explored the 4 split functions

[33mcommit 6cac6800c8a3d74b88089bb193d60e6708e5b0a9[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 22:41:31 2023 +0630

    Added explore again. Explored Split briefly

[33mcommit 940a1f1c5444957b67fe00b7c8e80a7e830c85d6[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 19:56:02 2023 +0630

    Reintroduced basics and put Join code in its own file

[33mcommit 2da545cc9dd8013fb5204f6f9ad026e54b17dce5[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 19:50:23 2023 +0630

    Revert "Added Join"
    
    This reverts commit 58487b2cfe289f5a3a55fa03d6c44a5055938f0f.

[33mcommit 58487b2cfe289f5a3a55fa03d6c44a5055938f0f[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 19:46:29 2023 +0630

    Added Join

[33mcommit fe68d0bdd8158f443ae3237ffb5cf0d1eb3ef416[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Sat Jun 3 15:21:48 2023 +0630

    Added explore and made most other code legacy

[33mcommit 4e319d5f2ce982f4e69b56dd3b03549aa7a16e31[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Fri Jun 2 23:37:45 2023 +0630

    included helpful comments for future development

[33mcommit e764306f5cde32e04437ae61ccdd248685ab8129[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Fri Jun 2 23:19:26 2023 +0630

    added further showcase in strings.go

[33mcommit 3a38b971278ef50d8ffd161074999fa34cf97d8f[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Fri Jun 2 23:03:07 2023 +0630

    furthered the strings.go file

[33mcommit aa89c7e1595df1f1ae8a24d75b4a6c4dbfc31a33[m
Merge: 361bafe 33cd015
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Fri Jun 2 22:17:37 2023 +0630

    Merge branch 'main' of https://github.com/DreamLineLove/gostrings
    Merge local development with newly created github repo.

[33mcommit 361bafe6086e5acaab64cdc6d009aebbf89609a5[m
Author: DreamLineLove <zwenyanzaw@protonmail.com>
Date:   Fri Jun 2 22:14:53 2023 +0630

    first commit

[33mcommit 33cd0154a78e758fcb1a5ba8564370ad32d2b317[m
Author: Zwe Nyan Zaw <zwenyanzaw@protonmail.com>
Date:   Fri Jun 2 22:11:37 2023 +0630

    Initial commit
