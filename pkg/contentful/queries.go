package contentful

const ProjectsQuery = `{ 
	projectCollection(order: sys_firstPublishedAt_DESC) {
		items {
			sys {
				id
				publishedAt
				firstPublishedAt
			}
			title
			description
			link
			thumbnail {
				url
				width
				height
				description
			}
			stack
		}
	}
}
`

const PostsQuery = `{
  postCollection(order: sys_firstPublishedAt_DESC) {
    items {
     sys {
				id
				publishedAt
				firstPublishedAt
			}
      title
      description
      thumbnail {
				url
				width
				height
				description
			}
      slug
      featured
      content
    }
  }
}
`
