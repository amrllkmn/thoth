<script lang="ts">
	import { page } from '$app/state';
</script>

<div class="mb-4 flex flex-col items-center gap-2 sm:mb-8">
	<h1 class="text-center text-2xl font-bold sm:text-3xl">{page.error?.message} 🫠</h1>
	<p class="text-center text-sm text-muted-foreground sm:text-base">
		Terribly sorry for this experience. Please contact the site owner.
	</p>
</div>
