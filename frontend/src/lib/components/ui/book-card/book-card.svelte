<script lang="ts">
	import * as Card from '../card';
	interface IBook {
		id: number;
		isnb13: string;
		isbn10: string;
		title: string;
		subtitle: string;
		authors: string;
		categories: string;
		thumbnail: string;
		description: string;
		published_year: number;
		average_rating: number;
		num_pages: number;
		ratings_count: number;
	}

	let { book }: { book: IBook } = $props();
</script>

<Card.Root class="h-full overflow-hidden transition-all hover:shadow-md">
	<Card.Header class="p-0">
		<div class="relative h-[150px] sm:h-[200px] w-full">
			<img
				src={book.thumbnail || '/placeholder.svg'}
				alt={`Cover of ${book.title}`}
				loading="lazy"
				decoding="async"
				class="absolute inset-0 h-full w-full object-cover"
			/>
		</div>
	</Card.Header>
  
</Card.Root>
