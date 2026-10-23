<script lang="ts">
  import { onMount, onDestroy } from 'svelte';
  import { Editor } from '@tiptap/core';
  import StarterKit from '@tiptap/starter-kit';
  import Link from '@tiptap/extension-link';
  import Image from '@tiptap/extension-image';
  import { 
    Bold, 
    Italic, 
    Heading2, 
    Heading3, 
    List, 
    ListOrdered, 
    Quote, 
    Link as LinkIcon, 
    Image as ImageIcon,
    Undo, 
    Redo 
  } from 'lucide-svelte';

  export let content = '';
  export let placeholder = 'Mulai menulis artikel Anda di sini...';

  let element: HTMLElement;
  let editor: Editor | null = null;

  onMount(() => {
    editor = new Editor({
      element: element,
      extensions: [
        StarterKit,
        Link.configure({
          openOnClick: false,
          HTMLAttributes: {
            class: 'text-brand-blue underline'
          }
        }),
        Image.configure({
          HTMLAttributes: {
            class: 'rounded-xl max-w-full my-4 shadow-sm'
          }
        })
      ],
      content: content,
      onUpdate: () => {
        if (editor) {
          content = editor.getHTML();
        }
      },
      editorProps: {
        attributes: {
          class: 'min-h-[280px] p-4 focus:outline-none text-ink-900 leading-relaxed font-body text-base'
        }
      }
    });
  });

  onDestroy(() => {
    if (editor) {
      editor.destroy();
    }
  });

  function addImage() {
    const url = window.prompt('Masukkan URL Gambar:');
    if (url && editor) {
      editor.chain().focus().setImage({ src: url }).run();
    }
  }

  function setLink() {
    if (!editor) return;
    const previousUrl = editor.getAttributes('link').href;
    const url = window.prompt('Masukkan URL:', previousUrl);

    if (url === null) return;
    if (url === '') {
      editor.chain().focus().extendMarkRange('link').unsetLink().run();
      return;
    }
    editor.chain().focus().extendMarkRange('link').setLink({ href: url }).run();
  }
</script>

<div class="rounded-xl border border-line-strong bg-white overflow-hidden shadow-sm">
  <!-- Toolbar -->
  {#if editor}
    <div class="flex flex-wrap items-center gap-1 border-b border-line-soft bg-surface-light p-2">
      <!-- Bold -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('bold') ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleBold().run()}
        title="Tebal"
      >
        <Bold size={16} />
      </button>

      <!-- Italic -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('italic') ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleItalic().run()}
        title="Miring"
      >
        <Italic size={16} />
      </button>

      <div class="h-4 w-px bg-line-soft mx-1"></div>

      <!-- Heading 2 -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('heading', { level: 2 }) ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleHeading({ level: 2 }).run()}
        title="Heading 2"
      >
        <Heading2 size={16} />
      </button>

      <!-- Heading 3 -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('heading', { level: 3 }) ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleHeading({ level: 3 }).run()}
        title="Heading 3"
      >
        <Heading3 size={16} />
      </button>

      <div class="h-4 w-px bg-line-soft mx-1"></div>

      <!-- Bullet List -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('bulletList') ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleBulletList().run()}
        title="Daftar Poin"
      >
        <List size={16} />
      </button>

      <!-- Ordered List -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('orderedList') ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleOrderedList().run()}
        title="Daftar Nomor"
      >
        <ListOrdered size={16} />
      </button>

      <!-- Quote -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('blockquote') ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={() => editor?.chain().focus().toggleBlockquote().run()}
        title="Kutipan"
      >
        <Quote size={16} />
      </button>

      <div class="h-4 w-px bg-line-soft mx-1"></div>

      <!-- Link -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors {editor.isActive('link') ? 'bg-white text-navy font-bold shadow-xs' : ''}"
        on:click={setLink}
        title="Tautan Link"
      >
        <LinkIcon size={16} />
      </button>

      <!-- Image -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors"
        on:click={addImage}
        title="Sisipkan Gambar"
      >
        <ImageIcon size={16} />
      </button>

      <div class="h-4 w-px bg-line-soft mx-1"></div>

      <!-- Undo / Redo -->
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors"
        on:click={() => editor?.chain().focus().undo().run()}
        title="Undo"
      >
        <Undo size={16} />
      </button>
      <button
        type="button"
        class="h-8 w-8 inline-flex items-center justify-center rounded text-ink-700 hover:bg-white transition-colors"
        on:click={() => editor?.chain().focus().redo().run()}
        title="Redo"
      >
        <Redo size={16} />
      </button>
    </div>
  {/if}

  <!-- Writing Area -->
  <div bind:this={element}></div>
</div>

<style>
  :global(.ProseMirror h2) {
    font-size: 1.5rem;
    font-weight: 700;
    margin-top: 1.5rem;
    margin-bottom: 0.75rem;
    color: var(--brand-navy);
  }
  :global(.ProseMirror h3) {
    font-size: 1.25rem;
    font-weight: 600;
    margin-top: 1.25rem;
    margin-bottom: 0.5rem;
    color: var(--brand-navy);
  }
  :global(.ProseMirror p) {
    margin-bottom: 0.75rem;
  }
  :global(.ProseMirror ul) {
    list-style-type: disc;
    padding-left: 1.5rem;
    margin-bottom: 0.75rem;
  }
  :global(.ProseMirror ol) {
    list-style-type: decimal;
    padding-left: 1.5rem;
    margin-bottom: 0.75rem;
  }
  :global(.ProseMirror blockquote) {
    border-left: 3px solid var(--accent-amber);
    padding-left: 1rem;
    font-style: italic;
    color: var(--ink-700);
    margin: 1rem 0;
  }
</style>
