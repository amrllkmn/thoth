
import { type TBook } from '$lib/types/types.js';

export async function load({ fetch }) {
  const res = await fetch('http://localhost:8080/v1/books?page=1&limit=10');
  const result = await res.json();

  console.log(result.books)


  return {
    books: result.books as TBook[],
  };
}