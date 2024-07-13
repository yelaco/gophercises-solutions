package parser

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseHTML(t *testing.T) {
	var tests = []struct {
		name  string
		input string
		want  []Link
	}{
		{"Only <a> tag", inputs[0], []Link{
			{
				Href: "/dog",
				Text: "Something in a span Text not in a span Bold text!",
			},
		}},
		{"No <a> tag", inputs[1], []Link{}},
		{"Single <a> tag", inputs[2], []Link{
			{
				Href: "https://www.javatpoint.com/html-favicon",
				Text: "this link",
			},
		}},
		{"Multiple <a> tags", inputs[3], []Link{
			{
				Href: "https://example.com/page1",
				Text: "Page 1",
			},
			{
				Href: "https://example.com/page2",
				Text: "Page 2",
			},
			{
				Href: "https://example.com/page3",
				Text: "Page 3",
			},
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			links, err := ParseHTML(tt.input)
			require.NoError(t, err)
			require.NotNil(t, links)
			require.Equal(t, len(tt.want), len(links))

			for i, link := range links {
				require.Equal(t, tt.want[i].Href, link.Href)
				require.Equal(t, tt.want[i].Text, link.Text)
			}
		})
	}
}

var inputs = []string{
	`
	<a href="/dog">
		<span>Something in a span</span>
		Text not in a span
		<b>Bold text!</b>
	</a>
	`,
	`
	<Html>
	<Head>
		<title>  
			Example of Paragraph tag  
		</title>
	</Head>
	<Body>
		<p>
			<!-- It is a Paragraph tag for creating the paragraph -->  
			<b> HTML </b> stands for <i> <u> Hyper Text Markup Language. </u> </i> It is used to create a web pages and applications. This language   
			is easily understandable by the user and also be modifiable. It is actually a Markup language, hence it provides a flexible way for designing the  
			web pages along with the text.   
		</p>
		HTML file is made up of different elements. <b> An element </b> is a collection of <i> start tag, end tag, attributes and the text between them</i>.   
		</p>  
	</Body>
	</Html>
	`,
	`
	<Html>
	<Head>
		<title>  
			Example of anchor or hyperlink  
		</title>
	</Head>
	<Body>
		<center> Click on <a href="https://www.javatpoint.com/html-favicon"> this link </a> for reading about HTML favicon in JavaTpoint.   </center>
	</Body>
	</Html>
	`,
	`
	<!DOCTYPE html>
	<html>
	<head>
		<title>Random HTML Example</title>
	</head>
	<body>
		<h1>Welcome to My Website</h1>
		<p>This is a sample HTML document with some random content and links.</p>
		
		<p>Check out these interesting links:</p>
		<ul>
			<li><a href="https://example.com/page1">Page 1</a></li>
			<li><a href="https://example.com/page2">Page 2</a></li>
			<li><a href="https://example.com/page3">Page 3</a></li>
		</ul>
		
		<p>Thank you for visiting!</p>
	</body>
	</html>
	`,
}
