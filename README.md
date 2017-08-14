# Marx

Another lightweight text formatter.

## Design

Marx is designed to achieve visual excellence when written in itself without rendering to some other format, and support basic notation for paragraph writing of Web text.

### Block Nodes

#### Heading

Headings are denoted by `@` and the number of the character indicates level of them.

```
@ Heading1
@@ Heading2
@@@ Heading3
```

Above will be converted into below:

```
<h1>Heading1</h1>
<h2>Heading2</h2>
<h3>Heading3</h3>
```

#### Paragraph

Marx only supports paragraph as a block of sentences. That is, it doesn't support line breaking signified by `<br>` tag in HTML. Which is because line breaking is unnecessary for paragraph writing.

```
Vladimir Vladimirovich Nabokov was a Russian-American novelist and entomologist. His first nine novels were in Russian, but he achieved international prominence after he began writing English prose.

Nabokov's Lolita (1955), his most noted novel in English, was ranked fourth in the list of the Modern Library 100 Best Novels; Pale Fire (1962) was ranked 53rd on the same list, and his memoir, Speak, Memory (1951), was listed eighth on the publisher's list of the 20th century's greatest nonfiction.
```

Above will be converted into below:

```
<p>Vladimir Vladimirovich Nabokov was a Russian-American novelist and entomologist. His first nine novels were in Russian, but he achieved international prominence after he began writing English prose.</p>

<p>Nabokov's Lolita (1955), his most noted novel in English, was ranked fourth in the list of the Modern Library 100 Best Novels; Pale Fire (1962) was ranked 53rd on the same list, and his memoir, Speak, Memory (1951), was listed eighth on the publisher's list of the 20th century's greatest nonfiction.</p>
```

#### Ordered List

An ordered list starts with `+` with arbitrary proceeding spaces. Indentation indicates hierarchy of the list.

```
+ foo
  + bar
  + baz
    + piyo
  + hoge
+ fuga
```

Above will be converted into below:

```
<ol>
  <li>foo</li>
  <li>
    <ol>
      <li>bar</li>
      <li>baz</li>
      <li>
        <ol>
          <li>piyo</li>
        </ol>
      </li>
    </ol>
  </li>
  <li>fuha</li>
</ol>
```

#### Unordered List

An ordered list starts with `-` with arbitrary proceeding spaces. Indentation indicates hierarchy of the list.

```
- foo
  - bar
  - baz
    - piyo
  - hoge
- fuga
```

Above will be converted into below:

```
<ul>
  <li>foo</li>
  <li>
    <ul>
      <li>bar</li>
      <li>baz</li>
      <li>
        <ul>
          <li>piyo</li>
        </ul>
      </li>
    </ul>
  </li>
  <li>fuha</li>
</ul>
```

### Inline Nodes

#### URL

```
https://example.com/
```

Above will be converted into below:

```
<a href="https://example.com/">https://example.com/</a>
```

#### URL with Options

URL can be annotated with some options.

##### `image`

```
https://example.com/example.png:image
```

Above will be converted into below:

```
<a href="https://example.com/example.png"><img src="https://example.com/example.png"></a>
```

##### `title`

```
https://example.com/:title="example page"
```

Above will be converted into below:

```
<a href="https://example.com/" title="example page">example page</a>
```

##### `image` with `title`

```
https://example.com/example.png:image:title="example image"
```

Above will be converted into below:

```
<a href="https://example.com/example.png" title="example image"><img src="https://example.com/example.png" alt="example image"></a>
```

#### Emphasis

```
**what you'd like to insist**
```

Above will be converted into below:

```
<em>what you'd like to insist</em>
```

#### Italic

```
/Oxford English Dictionary/
```

Above will be converted into below:

```
<i>Oxford English Dictionary</em>
```

## Author

Kentaro Kuribayashi <kentarok@gmail.com>

## License

MIT