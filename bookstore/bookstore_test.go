package bookstore_test

import (
	"bookstore"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		Title:  "For the Love of Go",
		Author: "John Ardundel",
		Copies: 9,
	}

	want := 8
	result, err := bookstore.Buy(book)

	if err != nil {
		t.Fatal(err)
	}

	got := result.Copies
	if got != want {
		t.Errorf("Failed, expected %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{}

	_, err := bookstore.Buy(b)

	if err == nil {
		t.Error("Want error buying from zero copies, got nil.")
	}
}

func TestGetAllBooks(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	want := []bookstore.Book{
		{ID: 1, Title: "For the Love of Go"},
		{ID: 2, Title: "The Power of Go: Tools"},
	}

	got := catalog.GetAllBooks()
	sort.Slice(got, func(i, j int) bool {
		return got[i].ID < got[j].ID
	})

	if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestGetBook(t *testing.T) {
	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go", ID: 1},
		2: {Title: "The Power of Go: Tools", ID: 2},
	}
	catalog[3] = bookstore.Book{ID: 3, Title: "Learning Go: An Idiomatic Approach"}

	for ID, book := range catalog {
		want := book

		got, err := catalog.GetBook(ID)

		if err != nil {
			t.Fatal(err)
		}

		if !cmp.Equal(want, got, cmpopts.IgnoreUnexported(bookstore.Book{})) {
			t.Error(cmp.Diff(want, got))
		}
	}

}

func TestGetInvalidBook(t *testing.T) {
	t.Parallel()

	catalog := bookstore.Catalog{
		1: {Title: "For the Love of Go", ID: 1},
		2: {Title: "The Power of Go: Tools", ID: 2},
	}
	catalog[3] = bookstore.Book{ID: 3, Title: "Learning Go: An Idiomatic Approach"}

	for ID := range catalog {
		got, err := catalog.GetBook(ID * 8)

		if err == nil {
			t.Fatalf("Expected error for books with invalid ID, got %v", got)
		}
	}
}

func TestNetPriceCents(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		PriceCents:      4000,
		DiscountPercent: 25,
	}

	want := 3000

	got := book.NetPriceCents()

	if want != got {
		t.Errorf("Expected %d. got %d in cents\n", want, got)
	}
}

func TestSetPriceCents(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		Title:      "Are we Web Yet",
		PriceCents: 450,
	}

	want := 5000

	err := book.SetPriceCents(5000)
	got := book.PriceCents

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func TestSetPriceCentsInvalid(t *testing.T) {
	t.Parallel()
	book := bookstore.Book{
		Title:      "Are we Web Yet",
		PriceCents: 450,
	}

	err := book.SetPriceCents(-4)

	if err == nil {
		t.Fatal("Expected an error for attempting to set invalid price, got nil.")
	}
}

func TestSetCategory(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		Title: "The Special One",
	}

	cats := []bookstore.Category{
		bookstore.CategoryAutobiography,
		bookstore.CategoryLargePrintRomance,
		bookstore.CategoryParticlePhysics,
	}

	for _, cat := range cats {
		want := cat

		err := book.SetCategory(cat)
		got := book.Category()

		if err != nil {
			t.Fatal(err)
		}

		if want != got {
			t.Errorf("Expected category to be %q, got %q", want, got)
		}
	}
}

func TestSetCategoryInvalid(t *testing.T) {
	t.Parallel()

	book := bookstore.Book{
		Title: "The Special One",
	}

	err := book.SetCategory(bookstore.Category(55))

	if err == nil {
		t.Error("Expected an error value for invalid category, got nil.")
	}
}
