---
title: DocuDex POC
description: A one-page document for the proof of concept — questions, appetite, what was built, what was found, recommendations — with the tests for each part
createdAt: 2026-09-26
updatedAt: 2026-09-29
---

1. **Recommendation.** *Given what you found, what should the project do?*
   Build differently. Fetch step becomes module cache + `go/doc` and unless the `go.mod` has drifted 
   DocuDex doesn't need an internet connection. If `go.mod` has drifted, DocuDex will run `go mod download`.
   The fetch pipeline also changes to reading downloaded zip files instead of the unpacked files so:
   - Zip is the version of record and the one that Go verifies against `go.sum`
   - Zip can be downloaded directly from the proxy so users can use DocuDex in their CI to generate and deploy
	 internal documentation sites.
   - There's useful metadata next to the zip files.
   DocuDex should also get the README from `GOMODCACHE` and include it at the top of the documentation.

2. **Question.** *What were you trying to find out?* 
   Does `go/doc` and related packages be used for getting useful API references, function signatures, 
   and examples for every dependency in a real `go.mod`? Yes!

3. **Appetite.** *How long was it worth, and how long did it take?* 
   2 hours budgeted. It took me a little over 2 hours since I was implementing everything by hand. 
   I did that because I don't have any experience with these packages. I got stuck in a few places and 
   got errors that I didn't know how to fix.
   
4. **What was built.** *What does the evidence come from?* 
   I wrote a few functions in my `main.go` that I built and ran inside Ink. My POC parsed `go.mod` 
   line by line, read the source files for those deps from `GOMODCACHE`, and used `go/doc` to generate 
   documentation and save those as markdown. The markdown doesn't have any sort of headings, links, 
   etc but I was able to get all API reference docs. The stuff that was missing was the docs from READMEs, etc.
   
5. **What was found.** *What actually happened, in numbers and names?*
   I got all the constants, variables, functions, and types. Some types had constants, variables, functions, 
   and methods. I read those too. All packages had their `README.md` at the top of the page on `pkg.go.dev`. 
   Those were not fetched since I only read `.go` files.

   Docs were generated for Ink and Temporal. Time was measured for benchmarking:
   Ink: 0.164973s, 6 direct dependencies, 68kb of markdown
   Temporal: 0.587530s, 81 direct dependencies, only 58 markdown files due to issues in implementation, 1.8mb of markdown 
   
   Surprise:
   I did go through some of the built-in features and I was surprise by how much was already available. I don't
   think I have to use any third-party packages for the core functionality - parsing, fetching and saving.
   
6. **Rabbit holes and next.** *What did this expose that the real project
   has to decide, and what is the next step?*
   - What is the exact pipeline for parsing `go.mod` to saving usable markdown?
   - What is the standardized markdown template and how do I get the godoc's markdown result to adhere to my template?
   - How to get the guides from READMEs, etc? If DocuDex had less documentation that `pkg.go.dev` or had a poorer DX, what's the point?
   - How to get agents to look at the local copies and not fetch from the internet?
   
   Next steps: 
   - Read the vanitydoc codebase and understand their documentation pipeline
   - Update my approach and get a minimum-viable CLI without markdown templates

References (for POC write-up template):
- [Singer, Write the pitch](https://basecamp.com/shapeup/1.5-chapter-06)
- [RFD 1](https://rfd.shared.oxide.computer/rfd/0001)
- [Design Docs at Google](https://www.industrialempathy.com/posts/design-docs-at-google/)
