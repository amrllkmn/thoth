import { API_URL } from '$env/static/private';
import { json } from '@sveltejs/kit';

export async function POST({ request }) {
	const { engine, query } = await request.json();

	const response = await fetch(`${API_URL}/v1/books/${engine}/search`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json'
		},
		body: JSON.stringify({
			query,
			pageInt: 1,
			limitInt: 10
		})
	});

	const result = await response.json();
	return json({ books: result.books, metadata: result.metadata }, { status: 200 });
}
