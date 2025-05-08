<!-- $lib/components/ui/loading-spinner/loading-spinner.svelte -->
<script>
	import { LoaderCircle } from 'lucide-svelte';

	let { size = 'lg', color = 'primary', strokeWidth = 2 } = $props();

	// Define size values in pixels for the Lucide icon
	const sizeValues = $derived(
		{
			sm: 24,
			md: 32,
			lg: 48,
			xl: 64
		}[size]
	);

	// Define color classes for the icon
	const colorClasses = $derived(
		{
			primary: 'text-primary',
			secondary: 'text-secondary',
			accent: 'text-accent',
			default: 'text-slate-800'
		}[color]
	);
</script>

<div class="flex items-center justify-center py-6">
	<div class={`animate-spin ${colorClasses}`} aria-label="Loading" role="status">
		<LoaderCircle size={sizeValues} {strokeWidth} />
	</div>
</div>
