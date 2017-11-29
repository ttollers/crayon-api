const chapterDir = process.env.CHAPTER_DIRECTORY || '/Users/tom/Projects/personal/book-lib/harry-potter-1/fr/chapter-1';
const native = require(`${ chapterDir }/native_final.json`);
const translated = require(`${ chapterDir }/translated_final.json`);

module.exports = {
  native: {
    get: () => native,
  },
  translated: {
    get: () => translated,
  },
};


