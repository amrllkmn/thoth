<script lang="ts">
	import { BookOpen } from 'lucide-svelte';

	let { logoClass = '' } = $props();
</script>

<div class={`relative ${logoClass}`}>
	<div class="absolute inset-0 flex items-center justify-center">
		<BookOpen stroke-width="1.5" class="h-full w-full" />
	</div>
</div>
