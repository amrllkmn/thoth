<script>
	import { BookCard } from '$lib/components/ui/book-card';
	import EmptyPage from '$lib/components/ui/empty-page/empty-page.svelte';
	import { EngineToggle } from '$lib/components/ui/engine-toggle';
	import { SearchBar } from '$lib/components/ui/search-bar';
	import { Title } from '$lib/components/ui/title';

	let { data } = $props();
	let query = $state('');
	let checked = $state(false);
	let searching = $state(false);

	// @ts-ignore
	async function handleSearch(event) {
		event.preventDefault();
		searching = true;
		const formData = new FormData(event.target);
		const query = formData.get('query') || '';
		const engine = formData.get('engine') || 'sqlite';

		const res = await fetch('/', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({
				query,
				engine
			})
		});

		const result = await res.json();
		data = { ...data, books: [...result.books] };
		searching = false;
	}
</script>

<Title />
<div class="space-y-4 sm:space-y-6">
	<form onsubmit={handleSearch} class="relative space-y-4 sm:space-y-6">
		<SearchBar {query} />
		<EngineToggle {checked} />
	</form>
	{#if searching}
		<p>Searching...</p>
	{:else if !searching && (!data.books || data.books.length < 1)}
		<EmptyPage />
	{:else}
		<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 sm:gap-6 lg:grid-cols-3">
			{#each data.books as book}
				<BookCard {book} />
			{/each}
		</div>
	{/if}
</div>
