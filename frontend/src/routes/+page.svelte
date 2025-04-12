<script>
	import { BookCard } from '$lib/components/ui/book-card';
	import { BookSearch } from '$lib/components/ui/book-search';
	import EmptyPage from '$lib/components/ui/empty-page/empty-page.svelte';
	import { EngineToggle } from '$lib/components/ui/engine-toggle';
	import { SearchBar } from '$lib/components/ui/search-bar';
	import { Title } from '$lib/components/ui/title';
	let books = [
		{
			id: 11,
			isbn13: '9780006479673',
			isbn10: '6479677',
			title: 'If Tomorrow Comes',
			subtitle: '',
			authors: 'Sidney Sheldon',
			categories: 'Adventure stories',
			thumbnail:
				'http://books.google.com/books/content?id=l2tBi_jLuk8C&printsec=frontcover&img=1&zoom=1&source=gbs_api',
			description:
				"One of Sidney Sheldon's most popular and bestselling titles, repackaged and reissued for a new generation of fans.",
			published_year: 1994,
			average_rating: 4.04,
			num_pages: 501,
			ratings_count: 49170
		},
		{
			id: 12,
			isbn13: '9780006480099',
			isbn10: '6480098',
			title: "Assassin's Apprentice",
			subtitle: '',
			authors: 'Robin Hobb',
			categories: 'American fiction',
			thumbnail:
				'http://books.google.com/books/content?id=qTaGQgAACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api',
			description: 'Fantasy-roman.',
			published_year: 1996,
			average_rating: 4.15,
			num_pages: 460,
			ratings_count: 133972
		},
		{
			id: 13,
			isbn13: '9780006482079',
			isbn10: '6482074',
			title: 'Warhost of Vastmark',
			subtitle: '',
			authors: 'Janny Wurts',
			categories: 'Fiction',
			thumbnail:
				'http://books.google.com/books/content?id=uOL0fpS9WZkC&printsec=frontcover&img=1&zoom=1&source=gbs_api',
			description:
				"Tricked once more by his wily half-brother, Lysaer, Lord of Light, arrives at the tiny harbor town of Merior to find that Arithon's ship yards have been abandoned and meticulously destroyed, and that the Master of Shadow has disappeared as if into thin air. Meanwhile Arithon and the Mad Prophet Dakar are traveling on foot through the treacherous Kelhorn Mountains towards the Vastmark clans, there to raise further support for his cause. But raising a warhost is a costly business. Is it mere coincidence that Princess Talith—Lysaer's beautiful, headstrong wife—is taken captive and held for a vast ransom by a master brigand? The forces of light and shadow circle and feint, drawing ever closer to a huge conflict. And in the background the Fellowship of Seven Sorcerers and the Koriani Enchantresses watch and plan, and wait.",
			published_year: 1995,
			average_rating: 4.03,
			num_pages: 522,
			ratings_count: 2966
		},
		{
			id: 14,
			isbn13: '9780006483014',
			isbn10: '6483011',
			title: 'The Once and Future King',
			subtitle: '',
			authors: 'Terence Hanbury White',
			categories: 'Arthurian romances',
			thumbnail:
				'http://books.google.com/books/content?id=Jx6BvgEACAAJ&printsec=frontcover&img=1&zoom=1&source=gbs_api',
			description:
				'An omnibus volume of the author\'s complete story of the Arthurian epic which includes: "The sword in the stone" (1939), "The witch in the wood" (1939), "The ill-made knight" (1940), "The candle in the wind" (published for the first time), and "The book of Merlyn." 1996.',
			published_year: 1996,
			average_rating: 4.04,
			num_pages: 823,
			ratings_count: 2805
		},
		{
			id: 15,
			isbn13: '9780006483892',
			isbn10: '6483895',
			title: 'Murder in LaMut',
			subtitle: '',
			authors: 'Raymond E. Feist;Joel Rosenberg',
			categories: 'Adventure stories',
			thumbnail:
				'http://books.google.com/books/content?id=I2jbBlMHlAMC&printsec=frontcover&img=1&zoom=1&source=gbs_api',
			description:
				'Available in the U.S. for the first time, this is the second volume in the exceptional Legends of the Riftwar series from "New York Times"-bestselling authors Feist and Rosenberg.',
			published_year: 2003,
			average_rating: 3.7,
			num_pages: 337,
			ratings_count: 5083
		}
	];
</script>

<Title />
<BookSearch />
<div class="space-y-4 sm:space-y-6">
	<SearchBar />
	<EngineToggle />
  {#if !books || books.length < 1}
    <EmptyPage/>
  {:else }
	<div class="grid grid-cols-1 gap-4 sm:grid-cols-2 sm:gap-6 lg:grid-cols-3">
		{#each books as book}
			<BookCard {book} />
		{/each}
	</div>
  {/if}
</div>
