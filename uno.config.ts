import { defineConfig, presetMini } from 'unocss'

export default defineConfig({
  presets: [presetMini({ dark: 'class' })],

  theme: {
    fontFamily: {
    },
    colors: {
      background: 'var(--bg)',
      foreground: 'var(--fg)',
      card: 'var(--card)',
      'card-foreground': 'var(--card-fg)',
      border: 'var(--border)',
      muted: 'var(--muted)',
      'muted-foreground': 'var(--muted-fg)',
      accent: 'var(--accent)',
      'accent-foreground': 'var(--accent-fg)',
      success: 'var(--success)',
    },
  },

  shortcuts: {
    'body-base': 'bg-background text-foreground font-sans antialiased',
    'nav-bar': 'fixed top-0 left-0 right-0 z-100 bg-background/80 border-b border-border',
    'nav-inner': 'max-w-3xl mx-auto px-6 h-14 flex items-center justify-between',
    'section-container': 'max-w-3xl mx-auto px-6 py-16',
    'section-heading': 'flex items-center gap-3 mb-8 font-mono text-accent',
    'section-heading-line': 'flex-1 h-px bg-border',
    'badge': 'inline-flex items-center px-2.5 py-1 rounded text-sm font-mono bg-card border border-border text-muted-foreground',
    'skill-card': 'p-4 rounded-lg border border-border bg-card',
    'card-base': 'rounded-lg border border-border bg-card p-6',
    'link': 'text-accent hover:underline underline-offset-4 decoration-accent/40',
    'timeline-dot': 'w-3 h-3 rounded-full bg-accent shrink-0 mt-1.5',
    'timeline-line': 'absolute left-1.5 top-6 bottom-0 w-px bg-border',
  },

  preflights: [
    {
      getCSS: () => `
        :root {
          --bg: #fafafa;
          --fg: #1a1a2e;
          --card: #ffffff;
          --card-fg: #1a1a2e;
          --border: #e5e7eb;
          --muted: #9ca3af;
          --muted-fg: #6b7280;
          --accent: #d97706;
          --accent-fg: #ffffff;
          --success: #059669;
          color-scheme: light;
        }

        .dark {
          --bg: #0d1117;
          --fg: #c9d1d9;
          --card: #161b22;
          --card-fg: #c9d1d9;
          --border: #30363d;
          --muted: #484f58;
          --muted-fg: #8b949e;
          --accent: #d2991d;
          --accent-fg: #0d1117;
          --success: #3fb950;
          color-scheme: dark;
        }

        * { box-sizing: border-box; margin: 0; padding: 0; }

        html {
          scroll-behavior: smooth;
          transition: color 0.2s, background-color 0.2s;
        }

        body {
          font-size: 1.0625rem;
          transition: color 0.2s, background-color 0.2s;
        }

        ::selection {
          background-color: var(--accent);
          color: var(--accent-fg);
        }

        [x-cloak] { display: none !important; }
      `,
    },
  ],
})
