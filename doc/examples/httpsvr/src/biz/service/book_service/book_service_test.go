/*
 * Copyright 2025 The Go-Spring Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package book_service

import (
	"testing"

	"httpsvr/src/dao/book_dao"

	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/gs/gstest"
	"github.com/lvan100/go-assert"
)

func init() {
	gstest.MockFor[*book_dao.BookDao]().With(&book_dao.BookDao{
		Store: map[string]book_dao.Book{
			"0": {SN: "0", Name: "Go Programing"},
		},
	})
	gs.Config().LocalFile.AddDir("../../../conf")
}

func TestMain(m *testing.M) {
	gstest.TestMain(m)
}

func TestBookService(t *testing.T) {

	x := gstest.Wire(t, new(struct {
		SvrAddr string            `value:"${server.addr}"`
		Service *BookService      `autowire:""`
		BookDao *book_dao.BookDao `autowire:""`
	}))

	assert.That(t, x.SvrAddr).Equal("0.0.0.0:9090")

	s, o := x.Service, x.BookDao
	assert.NotNil(t, o)

	books, err := s.ListBooks()
	assert.Nil(t, err)
	assert.That(t, len(books)).Equal(1)

	err = s.SaveBook(book_dao.Book{SN: "1", Name: "Go Spring"})
	assert.Nil(t, err)

	books, err = s.ListBooks()
	assert.Nil(t, err)
	assert.That(t, len(books)).Equal(2)
	assert.That(t, books[1].SN).Equal("1")
	assert.That(t, books[1].Name).Equal("Go Spring")

	book, err := s.GetBook("1")
	assert.Nil(t, err)
	assert.That(t, book.SN).Equal("1")
	assert.That(t, book.Name).Equal("Go Spring")

	err = s.DeleteBook("1")
	assert.Nil(t, err)

	books, err = s.ListBooks()
	assert.Nil(t, err)
	assert.That(t, len(books)).Equal(1)
}
