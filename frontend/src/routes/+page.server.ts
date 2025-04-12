
import { type TBook } from '$lib/types/types.js';
import { error as errorPage } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';

export async function load({ fetch }) {
  try {
    const res = await fetch(`${API_URL}/v1/books?page=1&limit=10`);
    const result = await res.json();
    
    if (res.status !== 200) {
      errorPage(400, 'Failed to fetch books');
    }
    return {
      books: result.books as TBook[],
    };
  } catch (error) {
    console.error('Error fetching books:', error);
    throw errorPage(500, 'Something went wrong')
  }

}