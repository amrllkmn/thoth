import { type TBook } from '$lib/types/types.js';
import { error as errorPage } from '@sveltejs/kit';
import { API_URL } from '$env/static/private';

export async function load({ fetch, url }) {
	try {
		const query = url.searchParams.get('query');
		const booksUrl = new URL(`${API_URL}/v1/books/search`);
		if (query && query.trim() !== '') {
			booksUrl.searchParams.set('query', query);
		}
		booksUrl.searchParams.set('page', '1');
		booksUrl.searchParams.set('limit', '10');
		const res = await fetch(booksUrl);
		const result = await res.json();

		if (res.status !== 200) {
			errorPage(400, 'Failed to fetch books');
		}
		return {
			books: result.books as TBook[]
		};
	} catch (error) {
		console.error('Error fetching books:', error);
		throw errorPage(500, 'Something went wrong');
	}
}
