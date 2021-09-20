import ultra from "https://raw.githubusercontent.com/exhibitionist-digital/ultra/master/mod.ts";

await ultra({
	importmap: await Deno.readTextFile("importmap.json"),
	lang: "en"
})
