<script lang="ts">
	import * as Card from '../card';
	import { Star } from 'lucide-svelte';
	interface IBook {
		id: number;
		isbn13: string;
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
	let primaryCategory = book.categories?.split(',')[0]?.trim() || 'Unknown';
	let formattedRating = book.average_rating ? book.average_rating.toFixed(1) : 'N/A';
</script>

<Card.Root class="h-full overflow-hidden transition-all hover:shadow-md">
	<Card.Header class="p-0">
		<div class="relative h-[150px] w-full sm:h-[200px]">
			<img
				src={book.thumbnail || '/placeholder.svg'}
				alt={`Cover of ${book.title}`}
				loading="lazy"
				decoding="async"
				class="absolute inset-0 h-full w-full object-cover"
			/>
			<div
				class="absolute left-2 top-2 rounded-md bg-muted-foreground px-2 py-1 text-xs text-secondary"
			>
				{primaryCategory}
			</div>
			<div
				class="absolute bottom-2 right-2 flex items-center gap-1 rounded-md bg-muted-foreground px-2 py-1 text-xs"
			>
				<Star class="text-gold-400 h-3 w-3 fill-secondary stroke-secondary" />
				<span class="text-secondary">{formattedRating}</span>
				<span class="text-[10px] text-base text-secondary">({book.ratings_count})</span>
			</div>
		</div>
	</Card.Header>
	<Card.Content class="p-3 sm:p-4">
		<h4 class="line-clamp-1 text-base font-bold sm:text-lg">
			{book.title}
		</h4>
		{#if book.subtitle}
			<p class="mb-1 line-clamp-1 text-xs text-muted-foreground">{book.subtitle}</p>
		{/if}
		<div class="mb-1 flex items-center justify-between text-xs sm:mb-2 sm:text-sm">
			<span class="line-clamp-1 flex-1 text-muted-foreground">{book.authors}</span>
			<span class="ml-2 whitespace-nowrap text-muted-foreground">{book.published_year}</span>
		</div>
		<p class="mb-2 line-clamp-2 text-xs text-muted-foreground sm:line-clamp-3 sm:text-sm">
			{book.description}
		</p>
	</Card.Content>
	<Card.Footer
		class="mt-auto flex items-center justify-between border-t pt-1 text-[10px] text-muted-foreground sm:text-xs"
	>
		<span>{book.num_pages} pages</span>
		<span class="max-w-[120px] truncate" title={`ISBN: ${book.isbn13}`}>
			ISBN: {book.isbn10 || book.isbn13}
		</span>
	</Card.Footer>
</Card.Root>
